module.exports = {
    testEnvironment: "jsdom",
    setupFilesAfterEnv: ["<rootDir>/jest.setup.js"],
    moduleNameMapper: {
        "\\.(css|less|scss|sass)$": "identity-obj-proxy", // 处理 CSS
        "^@/(.*)$": "<rootDir>/src/$1", // 路径别名
        "^@/gen/(.*)$": "<rootDir>/gen/$1" // 生成文件路径
    },
    transform: {
        "^.+\\.(js|jsx|ts|tsx)$": "babel-jest" // Babel 转换
    }
};