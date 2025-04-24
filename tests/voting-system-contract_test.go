package tests

import (
	"errors"
	"testing"

	"github.com/VikasFalcon/voting-system-chaincode/contracts"
	"github.com/VikasFalcon/voting-system-chaincode/mocks"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/stretchr/testify/require"
)

func TestIsStateExists(t *testing.T) {

	type args struct {
		key            string
		prepareMocking func(*mocks.FakeChaincodeStubInterface)
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "State exists",
			args: args{
				key: "voter1",
				prepareMocking: func(stub *mocks.FakeChaincodeStubInterface) {
					stub.GetStateStub = func(key string) ([]byte, error) {
						if key == "voter1" {
							return []byte(`some-value`), nil
						}
						return nil, nil
					}
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "State not found",
			args: args{
				key: "non-existing-key",
				prepareMocking: func(stub *mocks.FakeChaincodeStubInterface) {
					stub.GetStateStub = func(key string) ([]byte, error) {
						return nil, nil
					}
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Error while fetching data",
			args: args{
				key: "error-key",
				prepareMocking: func(stub *mocks.FakeChaincodeStubInterface) {
					stub.GetStateStub = func(key string) ([]byte, error) {
						return nil, errors.New("data not found error")
					}
				},
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := &mocks.FakeChaincodeStubInterface{}
			ctx := &mocks.FakeTransactionContextInterface{}

			ctx.GetStubStub = func() shim.ChaincodeStubInterface {
				return stub
			}

			tt.args.prepareMocking(stub)

			s := new(contracts.VotingContract)
			got, err := s.IsStateExists(ctx, tt.args.key)

			if (err != nil) != tt.wantErr {
				t.Errorf("IsStateExists(): error = %v, wantEr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.want, got)
		})
	}
}
