import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { unauthenticatedApi, authenticatedApi } from '../api';

interface User {
  id: string;
  name: string;
  email: string;
}

interface AuthContextType {
  user: User | null;
  isAuthenticated: boolean;
  loading: boolean;
  login: (email: string, password: string) => Promise<void>;
  signup: (name: string, email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  checkAuthStatus: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [loading, setLoading] = useState(true);

  // 初回マウント時に認証状態をチェック
  useEffect(() => {
    checkAuthStatus();
  }, []);

  // 認証状態を確認する
  const checkAuthStatus = async () => {
    try {
      setLoading(true);
      const response = await authenticatedApi('GET', 'user/index');
      
      if (response.status === 'success') {
        setUser(response.data);
        setIsAuthenticated(true);
      } else {
        setUser(null);
        setIsAuthenticated(false);
      }
    } catch (error) {
      console.error('認証チェックエラー:', error);
      setUser(null);
      setIsAuthenticated(false);
    } finally {
      setLoading(false);
    }
  };

  // ログイン
  const login = async (email: string, password: string) => {
    try {
      const response = await unauthenticatedApi('POST', 'user/signin', {
        email,
        password,
      });

      if (response.status === 'success') {
        setUser(response.data.user);
        setIsAuthenticated(true);
        return response.data;
      } else {
        throw new Error(response.message || 'ログインに失敗しました');
      }
    } catch (error) {
      console.error('ログインエラー:', error);
      throw error;
    }
  };

  // サインアップ
  const signup = async (name: string, email: string, password: string) => {
    try {
      const response = await unauthenticatedApi('POST', 'user/signup', {
        name,
        email,
        password,
      });

      if (response.status === 'success') {
        setUser(response.data.user);
        setIsAuthenticated(true);
        return response.data;
      } else {
        throw new Error(response.message || 'ユーザー登録に失敗しました');
      }
    } catch (error) {
      console.error('サインアップエラー:', error);
      throw error;
    }
  };

  // ログアウト
  const logout = async () => {
    try {
      await authenticatedApi('POST', 'user/logout');
    } catch (error) {
      console.error('ログアウトエラー:', error);
    } finally {
      setUser(null);
      setIsAuthenticated(false);
    }
  };

  const value = {
    user,
    isAuthenticated,
    loading,
    login,
    signup,
    logout,
    checkAuthStatus,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};