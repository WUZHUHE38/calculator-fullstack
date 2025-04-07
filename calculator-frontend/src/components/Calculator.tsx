"use client";
import React, {useState} from "react";
import {createClient} from "@connectrpc/connect";
import {createConnectTransport} from "@connectrpc/connect-web";
import {CalculatorService} from "@/gen/proto/calculator/v1/calculator_connect";

export default function Calculator() {
    const [operand1, setOperand1] = useState("");
    const [operand2, setOperand2] = useState("");
    const [operator, setOperator] = useState("+");
    const [result, setResult] = useState("");
    const [error, setError] = useState("");

    const transport = createConnectTransport({
        baseUrl: "http://localhost:8080", // 后端地址
    });

    const client = createClient(CalculatorService, transport);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const num1 = parseFloat(operand1);
        const num2 = parseFloat(operand2);

        if (isNaN(num1)) {
            setError("Operand 1 must be a valid number");
            return;
        }
        if (isNaN(num2)) {
            setError("Operand 2 must be a valid number");
            return;
        }

        try {
            const response = await client.calculate({
                operand1: num1,
                operand2: num2,
                operator: operator,
            });
            setResult(response.result.toString());
            setError("");
        } catch (err) {
            setError(err instanceof Error ? err.message : "Unknown error");
            setResult("");
        }
    };

    return (
        <div className="max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <input
                        type="number"
                        value={operand1}
                        onChange={(e) => setOperand1(e.target.value)}
                        placeholder="First number"
                        className="w-full p-2 border rounded"
                        required
                    />
                </div>
                <div>
                    <select
                        value={operator}
                        onChange={(e) => setOperator(e.target.value)}
                        className="w-full p-2 border rounded"
                    >
                        <option value="+">+</option>
                        <option value="-">-</option>
                        <option value="*">×</option>
                        <option value="/">÷</option>
                    </select>
                </div>
                <div>
                    <input
                        type="number"
                        value={operand2}
                        onChange={(e) => setOperand2(e.target.value)}
                        placeholder="Second number"
                        className="w-full p-2 border rounded"
                        required
                    />
                </div>
                <button
                    type="submit"
                    className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
                >
                    Calculate
                </button>
            </form>
            {error && (
                <div className="mt-4 p-2 bg-red-100 text-red-700 rounded" >{error}</div>
            )}
            {result && (
                <div className="mt-4 p-4 bg-green-100 text-green-700 rounded">
                    Result: {result}
                </div>
            )}
        </div>
    );
}