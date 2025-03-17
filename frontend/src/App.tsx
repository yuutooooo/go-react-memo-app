// src/App.js
import React, { useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import Top from "./pages/Top";
import Home from "./pages/Home";

function App() {
  // ログイン状態を管理するstate
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  // ログイン状態をチェックする関数
  const checkAuth = () => {
    // 通常はローカルストレージやCookieからトークンを取得して検証します
    // ここではusestateだけで簡易的に管理
    return isLoggedIn;
  };

  // 認証が必要なルートをラップするコンポーネント
  const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
    if (!checkAuth()) {
      // 認証されていない場合はログインページにリダイレクト
      return <Navigate to="/login" replace />;
    }
    return <>{children}</>;
  };

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Top />} />
        <Route
          path="/login"
          element={<Login setIsLoggedIn={setIsLoggedIn} />}
        />
        <Route path="/signup" element={<Signup />} />
        <Route path="/home" element={<Home />} />
        {/* 存在しないパスの場合はホームにリダイレクト */}
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </Router>
  );
}

export default App;
