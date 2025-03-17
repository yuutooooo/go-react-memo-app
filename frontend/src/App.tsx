// src/App.js
import React from 'react';
import Button from '@mui/material/Button';

function App() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="p-8 bg-white rounded shadow-lg">
        <h1 className="text-2xl font-bold mb-4">MUI と Tailwind の併用サンプル</h1>
        <Button variant="contained" color="primary">
          MUI Button
        </Button>
        <p className="mt-4 text-gray-600">
          これは Tailwind のユーティリティクラスを使用したテキストです。
        </p>
      </div>
    </div>
  );
}

export default App;
