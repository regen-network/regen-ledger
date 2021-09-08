package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/orm/testdata"
	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestPaginationProperty(t *testing.T) {
	t.Run("TestPagination", rapid.MakeCheck(func(t *rapid.T) {
		// Create a slice of group members
		groupMembers := rapid.SliceOf(genGroupMember).Draw(t, "groupMembers").([]*testdata.GroupMember)

		// Choose a random limit for paging
		upperLimit := uint64(len(groupMembers))
		if upperLimit == 0 {
			upperLimit = 1
		}
		limit := rapid.Uint64Range(1, upperLimit).Draw(t, "limit").(uint64)

		// Reconstruct the slice from offset pages
		reconstructedGroupMembers := make([]*testdata.GroupMember, 0, len(groupMembers))
		for offset := uint64(0); offset < uint64(len(groupMembers)); offset += limit {
			pageRequest := &query.PageRequest{
				Key:        nil,
				Offset:     offset,
				Limit:      limit,
				CountTotal: false,
				Reverse:    false,
			}
			end := offset + limit
			if end > uint64(len(groupMembers)) {
				end = uint64(len(groupMembers))
			}
			dest := reconstructedGroupMembers[offset:end]
			groupMembersIt := testGroupMemberIterator(groupMembers, nil)
			orm.Paginate(groupMembersIt, pageRequest, &dest)
			reconstructedGroupMembers = append(reconstructedGroupMembers, dest...)
		}

		// Should be the same slice
		require.Equal(t, len(groupMembers), len(reconstructedGroupMembers))
		for i, gm := range groupMembers {
			require.Equal(t, *gm, *reconstructedGroupMembers[i])
		}

		// Reconstruct the slice from keyed pages
		reconstructedGroupMembers = make([]*testdata.GroupMember, 0, len(groupMembers))
		var start uint64 = 0
		key := orm.EncodeSequence(0)
		for key != nil {
			pageRequest := &query.PageRequest{
				Key:        key,
				Offset:     0,
				Limit:      limit,
				CountTotal: false,
				Reverse:    false,
			}

			end := start + limit
			if end > uint64(len(groupMembers)) {
				end = uint64(len(groupMembers))
			}

			dest := reconstructedGroupMembers[start:end]
			groupMembersIt := testGroupMemberIterator(groupMembers, key)

			resp, err := orm.Paginate(groupMembersIt, pageRequest, &dest)
			require.NoError(t, err)
			key = resp.NextKey

			reconstructedGroupMembers = append(reconstructedGroupMembers, dest...)

			start += limit
		}

		// Should be the same slice
		require.Equal(t, len(groupMembers), len(reconstructedGroupMembers))
		for i, gm := range groupMembers {
			require.Equal(t, *gm, *reconstructedGroupMembers[i])
		}
	}))
}

func testGroupMemberIterator(gms []*testdata.GroupMember, key orm.RowID) orm.Iterator {
	var closed bool
	var index int
	if key != nil {
		index = int(orm.DecodeSequence(key))
	}
	return orm.IteratorFunc(func(dest codec.ProtoMarshaler) (orm.RowID, error) {
		if dest == nil {
			return nil, errors.Wrap(orm.ErrArgument, "destination object must not be nil")
		}

		if index == len(gms) {
			closed = true
		}

		if closed {
			return nil, orm.ErrIteratorDone
		}

		rowID := orm.EncodeSequence(uint64(index))

		bytes, err := gms[index].Marshal()
		if err != nil {
			return nil, err
		}

		index++

		return rowID, dest.Unmarshal(bytes)
	})
}
