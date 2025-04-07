import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import Calculator from "@/components/Calculator";
import { createClient } from "@connectrpc/connect";

// Mock ConnectRPC 客户端
jest.mock("@connectrpc/connect", () => ({
    createClient: jest.fn(() => ({
        calculate: jest.fn(),
    })),
    createConnectTransport: jest.fn(),
}));

// Mock 生成的 proto 服务
jest.mock("@/gen/proto/calculator/v1/calculator_connect", () => ({
    CalculatorService: { typeName: "CalculatorService" },
}));

describe("Calculator Component", () => {
    const mockCalculate = jest.fn();
    const mockClient = {
        calculate: mockCalculate,
    };

    beforeEach(() => {
        (createClient as jest.Mock).mockReturnValue(mockClient);
        mockCalculate.mockClear();
    });

    it("renders the calculator form", () => {
        render(<Calculator />);

        expect(screen.getByPlaceholderText("First number")).toBeInTheDocument();
        expect(screen.getByPlaceholderText("Second number")).toBeInTheDocument();
        expect(screen.getByRole("button", { name: "Calculate" })).toBeInTheDocument();
    });

    it("performs addition and displays result", async () => {
        mockCalculate.mockResolvedValueOnce({ result: 5 });
        render(<Calculator />);

        // 填写表单
        fireEvent.change(screen.getByPlaceholderText("First number"), {
            target: { value: "2" },
        });
        fireEvent.change(screen.getByPlaceholderText("Second number"), {
            target: { value: "3" },
        });
        fireEvent.click(screen.getByRole("button"));

        // 验证 API 调用
        await waitFor(() => {
            expect(mockCalculate).toHaveBeenCalledWith({
                operand1: 2,
                operand2: 3,
                operator: "+",
            });
        });

        // 验证结果显示
        expect(await screen.findByText("Result: 5")).toBeInTheDocument();
    });

    it("handles division by zero error", async () => {
        mockCalculate.mockRejectedValueOnce(new Error("division by zero"));
        render(<Calculator />);

        // 填写表单
        fireEvent.change(screen.getByPlaceholderText("First number"), {
            target: { value: "5" },
        });
        fireEvent.change(screen.getByPlaceholderText("Second number"), {
            target: { value: "0" },
        });
        fireEvent.change(screen.getByRole("combobox"), {
            target: { value: "/" },
        });
        fireEvent.click(screen.getByRole("button"));

        // 验证错误显示
        expect(await screen.findByText("division by zero")).toBeInTheDocument();
        expect(screen.queryByText(/Result:/)).not.toBeInTheDocument();
    });

    it("handles different operators", async () => {
        mockCalculate.mockResolvedValueOnce({ result: 6 });
        render(<Calculator />);

        // 修改运算符
        fireEvent.change(screen.getByRole("combobox"), {
            target: { value: "*" },
        });
        fireEvent.change(screen.getByPlaceholderText("First number"), {
            target: { value: "2" },
        });
        fireEvent.change(screen.getByPlaceholderText("Second number"), {
            target: { value: "3" },
        });
        fireEvent.click(screen.getByRole("button"));

        await waitFor(() => {
            expect(mockCalculate).toHaveBeenCalledWith(
                expect.objectContaining({ operator: "*" })
            );
        });
    });
});