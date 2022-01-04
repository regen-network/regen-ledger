package server

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"testing"
)

func Test_checkFilters(t *testing.T) {
	type args struct {
		filters     []*ecocredit.Filter
		classInfo   ecocredit.ClassInfo
		batchInfo   ecocredit.BatchInfo
		projectInfo ecocredit.ProjectInfo
		owner       string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid single depth filter",
			args: args{
				filters:     []*ecocredit.Filter{{Sum: &ecocredit.Filter_ProjectId{ProjectId: "P01"}}},
				classInfo:   ecocredit.ClassInfo{},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{ProjectId: "P01"},
				owner:       "",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "valid AND filter",
			args: args{
				filters: []*ecocredit.Filter{
					{Sum: generateANDFilter([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}},
						{Sum: &ecocredit.Filter_CreditTypeName{CreditTypeName: "BRUH"}}})}},
				classInfo:   ecocredit.ClassInfo{ClassId: "F00", CreditType: &ecocredit.CreditType{Name: "BRUH"}},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "valid OR filter (only needs 1 matcher)",
			args: args{
				filters: []*ecocredit.Filter{
					{Sum: generateORFilters([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}},
						{Sum: &ecocredit.Filter_CreditTypeName{CreditTypeName: "BRUH"}}})}},
				classInfo: ecocredit.ClassInfo{ClassId: "NOTF00", CreditType: &ecocredit.CreditType{
					Name:         "BRUH",
					Abbreviation: "",
					Unit:         "",
					Precision:    0,
				}},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "invalid single depth filter",
			args: args{
				filters:     []*ecocredit.Filter{{Sum: &ecocredit.Filter_ClassId{ClassId: "BAD"}}},
				classInfo:   ecocredit.ClassInfo{ClassId: "OOPS"},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			want:    1,
			wantErr: true,
		},
		{
			name: "invalid AND filter -- 1 non-match",
			args: args{
				filters: []*ecocredit.Filter{
					{Sum: generateANDFilter([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}},
						{Sum: &ecocredit.Filter_ProjectId{ProjectId: "B4Z"}}})}},
				classInfo:   ecocredit.ClassInfo{ClassId: "F00"},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{ProjectId: "NOPE"},
				owner:       "",
			},
			want:    1,
			wantErr: true,
		},
		{
			name: "invalid OR filter -- no matches",
			args: args{
				filters: []*ecocredit.Filter{
					{Sum: generateORFilters([]*ecocredit.Filter{
						{Sum: &ecocredit.Filter_ProjectId{ProjectId: "B4Z"}},
						{Sum: &ecocredit.Filter_ClassId{ClassId: "F00"}}})}},
				classInfo:   ecocredit.ClassInfo{ClassId: "BAD"},
				batchInfo:   ecocredit.BatchInfo{},
				projectInfo: ecocredit.ProjectInfo{ProjectId: "BAD2"},
				owner:       "",
			},
			want:    2,
			wantErr: true,
		},
		{
			name: "valid multi-depth OR filter",
			args: args{
				// FILTER LAYOUT
				// a V b V c
				//         |
				//       c V d
				filters: []*ecocredit.Filter{
					{Sum: generateORFilters([]*ecocredit.Filter{ // OR
						{Sum: &ecocredit.Filter_ProjectId{ProjectId: "F00"}},
						{Sum: &ecocredit.Filter_ClassAdmin{ClassAdmin: "ME"}},
						{Sum: generateORFilters([]*ecocredit.Filter{ // OR
							{Sum: &ecocredit.Filter_ProjectLocation{ProjectLocation: "NOWHERE"}},
							{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "YES"}}})}})}},
				classInfo:   ecocredit.ClassInfo{},
				batchInfo:   ecocredit.BatchInfo{BatchDenom: "YES"},
				projectInfo: ecocredit.ProjectInfo{},
				owner:       "",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "valid OR nested in AND",
			args: args{
				// filter layout:
				// a ^ b ^ c
				//         |
				//       d v e
				filters: []*ecocredit.Filter{
					{Sum: generateANDFilter([]*ecocredit.Filter{ // AND
						{Sum: &ecocredit.Filter_ClassId{ClassId: "ME"}},
						{Sum: &ecocredit.Filter_ProjectLocation{ProjectLocation: "HERE"}},
						{Sum: generateORFilters([]*ecocredit.Filter{ // OR
							{Sum: &ecocredit.Filter_ProjectId{ProjectId: "YES"}},
							{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "DENOM"}}},
						)}},
					)}},
				classInfo:   ecocredit.ClassInfo{ClassId: "ME"},
				batchInfo:   ecocredit.BatchInfo{BatchDenom: "DENOM"},
				projectInfo: ecocredit.ProjectInfo{ProjectLocation: "HERE"},
				owner:       "",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkFilters(tt.args.filters, tt.args.classInfo, tt.args.batchInfo, tt.args.projectInfo, tt.args.owner)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFilters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkFilters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateANDFilter(filters []*ecocredit.Filter) *ecocredit.Filter_And_ {
	return &ecocredit.Filter_And_{And: &ecocredit.Filter_And{Filters: filters}}
}

func generateORFilters(filters []*ecocredit.Filter) *ecocredit.Filter_Or_ {
	return &ecocredit.Filter_Or_{Or: &ecocredit.Filter_Or{Filters: filters}}
}
