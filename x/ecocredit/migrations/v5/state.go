package v5

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
)

// MigrateState performs in-place store migrations from ConsensusVersion 4 to 5.
func MigrateState(sdkCtx sdk.Context, baseStore ecocreditv1.StateStore) error {
	// collect all the classes associated with batches by looking up the class key from the project
	batchIt, err := baseStore.BatchTable().List(sdkCtx, ecocreditv1.BatchPrimaryKey{})
	if err != nil {
		return err
	}

	batchClasses := map[uint64]uint64{}
	for batchIt.Next() {
		batch, err := batchIt.Value()
		if err != nil {
			return err
		}

		if batch.ClassKey != 0 {
			return fmt.Errorf("unexpected state, expected batch class key to be 0 before migration, got %d", batch.ClassKey)
		}

		proj, err := baseStore.ProjectTable().Get(sdkCtx, batch.ProjectKey)
		if err != nil {
			return err
		}

		batchClasses[batch.Key] = proj.ClassKey
	}
	batchIt.Close()

	// set class keys on batches
	for batchKey, classKey := range batchClasses {
		batch, err := baseStore.BatchTable().Get(sdkCtx, batchKey)
		if err != nil {
			return err
		}

		batch.ClassKey = classKey
		if err := baseStore.BatchTable().Save(sdkCtx, batch); err != nil {
			return err
		}
	}

	// collect all the classes associated with projects
	projectIt, err := baseStore.ProjectTable().List(sdkCtx, ecocreditv1.ProjectPrimaryKey{})
	if err != nil {
		return err
	}

	projectClasses := map[uint64]uint64{}
	for projectIt.Next() {
		proj, err := projectIt.Value()
		if err != nil {
			return err
		}

		if proj.ClassKey == 0 {
			return fmt.Errorf("unexpected state, expected project class key to be non-zero before migration, got %d", proj.ClassKey)
		}

		projectClasses[proj.Key] = proj.ClassKey
	}
	projectIt.Close()

	// create enrollment entries for all project class relationships
	for projectKey, classKey := range projectClasses {
		err = baseStore.ProjectEnrollmentTable().Insert(sdkCtx, &ecocreditv1.ProjectEnrollment{
			ProjectKey: projectKey,
			ClassKey:   classKey,
			Status:     ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_ACCEPTED,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
