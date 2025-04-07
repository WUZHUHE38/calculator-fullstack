package service

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	calculatorv1 "github.com/WUZHUHE38/calculator-fullstack/calculator-backend/gen/go/calculator/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatorService_Calculate(t *testing.T) {
	service := &CalculatorService{}
	ctx := context.Background()

	testCases := []struct {
		name        string
		req         *calculatorv1.CalculateRequest
		wantResult  float64
		wantErrCode connect.Code
	}{
		{
			name: "addition",
			req: &calculatorv1.CalculateRequest{
				Operand1: 5,
				Operand2: 3,
				Operator: "+",
			},
			wantResult: 8,
		},
		{
			name: "subtraction",
			req: &calculatorv1.CalculateRequest{
				Operand1: 10,
				Operand2: 4,
				Operator: "-",
			},
			wantResult: 6,
		},
		{
			name: "multiplication",
			req: &calculatorv1.CalculateRequest{
				Operand1: 7,
				Operand2: 3,
				Operator: "*",
			},
			wantResult: 21,
		},
		{
			name: "division",
			req: &calculatorv1.CalculateRequest{
				Operand1: 15,
				Operand2: 5,
				Operator: "/",
			},
			wantResult: 3,
		},
		{
			name: "division by zero",
			req: &calculatorv1.CalculateRequest{
				Operand1: 10,
				Operand2: 0,
				Operator: "/",
			},
			wantErrCode: connect.CodeInvalidArgument,
		},
		{
			name: "invalid operator",
			req: &calculatorv1.CalculateRequest{
				Operand1: 5,
				Operand2: 2,
				Operator: "%",
			},
			wantErrCode: connect.CodeInvalidArgument,
		},
		{
			name: "decimal numbers",
			req: &calculatorv1.CalculateRequest{
				Operand1: 3.5,
				Operand2: 1.5,
				Operator: "+",
			},
			wantResult: 5.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			connectReq := connect.NewRequest(tc.req)
			resp, err := service.Calculate(ctx, connectReq)

			if tc.wantErrCode != 0 {
				require.Error(t, err)
				connectErr, ok := err.(*connect.Error)
				require.True(t, ok, "error should be a connect error")
				assert.Equal(t, tc.wantErrCode, connectErr.Code())
				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.wantResult, resp.Msg.Result)
		})
	}
}

func TestCalculatorService_EdgeCases(t *testing.T) {
	service := &CalculatorService{}
	ctx := context.Background()

	t.Run("zero addition", func(t *testing.T) {
		req := connect.NewRequest(&calculatorv1.CalculateRequest{
			Operand1: 0,
			Operand2: 0,
			Operator: "+",
		})
		resp, err := service.Calculate(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, 0.0, resp.Msg.Result)
	})

	t.Run("negative numbers", func(t *testing.T) {
		req := connect.NewRequest(&calculatorv1.CalculateRequest{
			Operand1: -5,
			Operand2: 3,
			Operator: "+",
		})
		resp, err := service.Calculate(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, -2.0, resp.Msg.Result)
	})

	t.Run("large numbers", func(t *testing.T) {
		req := connect.NewRequest(&calculatorv1.CalculateRequest{
			Operand1: 1e18,
			Operand2: 2e18,
			Operator: "+",
		})
		resp, err := service.Calculate(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, 3e18, resp.Msg.Result)
	})
}
