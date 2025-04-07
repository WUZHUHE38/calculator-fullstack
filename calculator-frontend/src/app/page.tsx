import Calculator from '@/components/Calculator';
import "./globals.css";

export default function Home() {
    return (
        <main className="min-h-screen p-24">
            <h1 className="text-4xl font-bold text-center mb-8">Calculator</h1>
            <Calculator />
        </main>
    )
}