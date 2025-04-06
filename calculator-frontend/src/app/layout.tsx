// src/app/layout.tsx
import "./globals.css";
import React from "react";  // 确保路径正确

export default function RootLayout({children,}: { children: React.ReactNode }) {
    return (
        <html lang="en">
        <body>{children}</body>
        </html>
    )
}