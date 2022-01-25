package server

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"testing"
)

func Test_checkFilters(t *testing.T) {
	type args struct {
		filter      *ecocredit.Filter
		classInfo   ecocredit.ClassInfo
		batchInfo   ecocredit.BatchInfo
		projectInfo ecocredit.ProjectInfo
		owner       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid single depth filter",
			args: args{
				filter:      &ecocredit.Filter{Sum: &ecocredit.Filter_ProjectId{ProjectId: "P01"}},
				classInfo:   ecocredit.ClassInfo{},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{ProjectId: "P01"},
				owner:       "",
			},
			wantErr: false,
		},
		{
			name: "valid AND filter",
			args: args{
				filter: &ecocredit.Filter{
					Sum: AND([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}},
						{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "BRUH"}}})},
				classInfo:   ecocredit.ClassInfo{ClassId: "F00"},
				batchInfo:   ecocredit.BatchInfo{BatchDenom: "BRUH"},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			wantErr: false,
		},
		{
			name: "valid OR filter (only needs 1 matcher)",
			args: args{
				filter: &ecocredit.Filter{
					Sum: OR([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}},
						{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "BRUH"}}})},
				classInfo:   ecocredit.ClassInfo{ClassId: "NOTF00"},
				batchInfo:   ecocredit.BatchInfo{BatchDenom: "BRUH"},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			wantErr: false,
		},
		{
			name: "invalid single depth filter",
			args: args{
				filter:      &ecocredit.Filter{Sum: &ecocredit.Filter_ClassId{ClassId: "BAD"}},
				classInfo:   ecocredit.ClassInfo{ClassId: "OOPS"},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			wantErr: true,
		},
		{
			name: "invalid AND filter -- 1 non-match",
			args: args{
				filter: &ecocredit.Filter{
					Sum: AND([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}},
						{Sum: &ecocredit.Filter_ProjectId{ProjectId: "B4Z"}}})},
				classInfo:   ecocredit.ClassInfo{ClassId: "F00"},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{ProjectId: "NOPE"},
				owner:       "",
			},
			wantErr: true,
		},
		{
			name: "invalid OR filter -- no matches",
			args: args{
				filter: &ecocredit.Filter{
					Sum: OR([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ProjectId{ProjectId: "B4Z"}},
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}}})},
				classInfo:   ecocredit.ClassInfo{ClassId: "BAD"},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{ProjectId: "BAD2"},
				owner:       "",
			},
			wantErr: true,
		},
		{
			name: "valid multi-depth OR filter",
			args: args{
				// FILTER LAYOUT
				// a V b V c
				//         |
				//       c V d
				filter: &ecocredit.Filter{
					Sum: OR([]*ecocredit.Filter{ // OR
						{Sum: &ecocredit.Filter_ProjectId{ProjectId: "F00"}},
						{Sum: &ecocredit.Filter_ClassAdmin{ClassAdmin: "ME"}},
						{Sum: OR([]*ecocredit.Filter{ // OR
							{Sum: &ecocredit.Filter_ProjectLocation{ProjectLocation: "NOWHERE"}},
							{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "YES"}}})}})},
				classInfo:   ecocredit.ClassInfo{},
				batchInfo:   ecocredit.BatchInfo{BatchDenom: "YES"},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			wantErr: false,
		},
		{
			name: "valid OR nested in AND",
			args: args{
				// filter layout:
				// a ^ b ^ c
				//         |
				//       d v e
				filter: &ecocredit.Filter{
					Sum: AND([]*ecocredit.Filter{ // AND
						{Sum: &ecocredit.Filter_ClassId{ClassId: "ME"}},
						{Sum: &ecocredit.Filter_ProjectLocation{ProjectLocation: "HERE"}},
						{Sum: OR([]*ecocredit.Filter{ // OR
							{Sum: &ecocredit.Filter_ProjectId{ProjectId: "YES"}},
							{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "DENOM"}}},
						)}},
					)},
				classInfo:   ecocredit.ClassInfo{ClassId: "ME"},
				batchInfo:   ecocredit.BatchInfo{BatchDenom: "DENOM"},
				projectInfo: ecocredit.ProjectInfo{ProjectLocation: "HERE"},
				owner:       "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkFilterMatch(tt.args.filter, tt.args.classInfo, tt.args.batchInfo, tt.args.projectInfo, tt.args.owner)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkCreditMatchesFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func AND(filters []*ecocredit.Filter) *ecocredit.Filter_And_ {
	return &ecocredit.Filter_And_{And: &ecocredit.Filter_And{Filters: filters}}
}

func OR(filters []*ecocredit.Filter) *ecocredit.Filter_Or_ {
	return &ecocredit.Filter_Or_{Or: &ecocredit.Filter_Or{Filters: filters}}
}
