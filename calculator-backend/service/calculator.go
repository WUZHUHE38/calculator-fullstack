package service

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	calculatorv1 "github.com/WUZHUHE38/calculator-fullstack/calculator-backend/gen/go/calculator/v1"
)

type CalculatorService struct{}

func (s *CalculatorService) Calculate(
	ctx context.Context,
	req *connect.Request[calculatorv1.CalculateRequest],
) (*connect.Response[calculatorv1.CalculateResponse], error) {
	op := req.Msg.Operator
	op1 := req.Msg.Operand1
	op2 := req.Msg.Operand2

	var result float64
	switch op {
	case "+":
		result = op1 + op2
	case "-":
		result = op1 - op2
	case "*":
		result = op1 * op2
	case "/":
		if op2 == 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("division by zero"))
		}
		result = op1 / op2
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid operator: %s", op))
	}
	return connect.NewResponse(&calculatorv1.CalculateResponse{Result: result}), nil
}
