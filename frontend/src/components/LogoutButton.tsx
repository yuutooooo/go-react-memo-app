import React from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, ButtonProps } from '@mui/material';
import { Logout as LogoutIcon } from '@mui/icons-material';
import { useAuth } from '../contexts/AuthContext';

interface LogoutButtonProps extends Omit<ButtonProps, 'onClick'> {
  showIcon?: boolean;
}

const LogoutButton: React.FC<LogoutButtonProps> = ({ 
  showIcon = true, 
  children, 
  ...props 
}) => {
  const navigate = useNavigate();
  const { logout } = useAuth();

  const handleLogout = async () => {
    try {
      await logout();
      navigate('/');
    } catch (error) {
      console.error('ログアウト中にエラーが発生しました', error);
    }
  };

  return (
    <Button
      onClick={handleLogout}
      startIcon={showIcon ? <LogoutIcon /> : undefined}
      {...props}
    >
      {children || 'ログアウト'}
    </Button>
  );
};

export default LogoutButton; 