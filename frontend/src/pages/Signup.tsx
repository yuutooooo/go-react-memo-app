import React, { useState } from "react";
import {
  Alert,
  Box,
  Button,
  Container,
  Paper,
  TextField,
  Typography,
  InputAdornment,
  IconButton,
  Divider,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import {
  Email as EmailIcon,
  Person as PersonIcon,
  Lock as LockIcon,
  Visibility as VisibilityIcon,
  VisibilityOff as VisibilityOffIcon,
  ArrowBack as ArrowBackIcon,
} from "@mui/icons-material";
import { Link, useNavigate } from "react-router-dom";

const Signup: React.FC = () => {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const navigate = useNavigate();
  const theme = useTheme();

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleClickShowConfirmPassword = () => {
    setShowConfirmPassword(!showConfirmPassword);
  };

  const handleSignup = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (
      email === "" ||
      username === "" ||
      password === "" ||
      confirmPassword === ""
    ) {
      setError("すべてのフィールドを入力してください");
      return;
    }
    
    if (password !== confirmPassword) {
      setError("パスワードが一致しません");
      return;
    }

    // サインアップ成功をシミュレート
    // 実際のアプリでは、APIに接続してアカウント作成を行います
    navigate("/login");
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-blue-50 to-white py-8 flex items-center justify-center">
      <Container maxWidth="sm">
        <Box className="mb-4 flex justify-center">
          <Button
            component={Link}
            to="/"
            startIcon={<ArrowBackIcon />}
            sx={{
              color: "#3f51b5",
              fontWeight: 500,
              "&:hover": {
                backgroundColor: "rgba(63, 81, 181, 0.04)",
              },
            }}
          >
            ホームに戻る
          </Button>
        </Box>

        <Paper
          elevation={0}
          sx={{
            borderRadius: "16px",
            border: "1px solid #e0e0ff",
            boxShadow: "0 4px 20px rgba(0,0,0,0.05)",
            overflow: "hidden",
          }}
        >
          <Box
            sx={{
              backgroundColor: "#3f51b5",
              py: 3,
              px: 2,
              color: "white",
              textAlign: "center",
            }}
          >
            <Typography variant="h5" component="h1" className="font-bold">
              新規アカウント登録
            </Typography>
          </Box>

          <Box sx={{ p: { xs: 3, md: 4 } }}>
            {error && (
              <Alert
                severity="error"
                className="mb-4"
                sx={{
                  borderRadius: "8px",
                  fontSize: "0.9rem",
                  alignItems: "center",
                }}
              >
                {error}
              </Alert>
            )}

            <form onSubmit={handleSignup}>
              <TextField
                label="メールアドレス"
                type="email"
                fullWidth
                margin="normal"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="例: user@example.com"
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <EmailIcon sx={{ color: "#3f51b5" }} />
                    </InputAdornment>
                  ),
                }}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    borderRadius: "8px",
                    "&:hover fieldset": {
                      borderColor: "#3f51b5",
                    },
                  },
                }}
              />

              <TextField
                label="ユーザー名"
                fullWidth
                margin="normal"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <PersonIcon sx={{ color: "#3f51b5" }} />
                    </InputAdornment>
                  ),
                }}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    borderRadius: "8px",
                    "&:hover fieldset": {
                      borderColor: "#3f51b5",
                    },
                  },
                }}
              />

              <TextField
                label="パスワード"
                type={showPassword ? "text" : "password"}
                fullWidth
                margin="normal"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <LockIcon sx={{ color: "#3f51b5" }} />
                    </InputAdornment>
                  ),
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        onClick={handleClickShowPassword}
                        edge="end"
                        size="small"
                      >
                        {showPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    borderRadius: "8px",
                    "&:hover fieldset": {
                      borderColor: "#3f51b5",
                    },
                  },
                }}
              />

              <TextField
                label="パスワード確認"
                type={showConfirmPassword ? "text" : "password"}
                fullWidth
                margin="normal"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <LockIcon sx={{ color: "#3f51b5" }} />
                    </InputAdornment>
                  ),
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        onClick={handleClickShowConfirmPassword}
                        edge="end"
                        size="small"
                      >
                        {showConfirmPassword ? (
                          <VisibilityOffIcon />
                        ) : (
                          <VisibilityIcon />
                        )}
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    borderRadius: "8px",
                    "&:hover fieldset": {
                      borderColor: "#3f51b5",
                    },
                  },
                }}
              />

              <Button
                type="submit"
                variant="contained"
                color="primary"
                fullWidth
                size="large"
                sx={{
                  mt: 4,
                  mb: 2,
                  borderRadius: "28px",
                  py: 1.2,
                  boxShadow: "0 4px 14px 0 rgba(0,118,255,0.39)",
                  backgroundColor: "#3f51b5",
                  fontSize: "1rem",
                  fontWeight: 600,
                  "&:hover": {
                    backgroundColor: "#303f9f",
                  },
                }}
              >
                アカウント作成
              </Button>
            </form>

            <Divider sx={{ my: 3 }}>
              <Typography variant="body2" color="textSecondary">
                または
              </Typography>
            </Divider>

            <Box sx={{ textAlign: "center" }}>
              <Typography
                variant="body1"
                className="mb-3"
                sx={{ fontWeight: 500 }}
              >
                すでにアカウントをお持ちの方は
              </Typography>
              <Button
                component={Link}
                to="/login"
                variant="outlined"
                fullWidth
                sx={{
                  borderRadius: "28px",
                  py: 1.2,
                  borderColor: "#3f51b5",
                  borderWidth: 2,
                  color: "#3f51b5",
                  fontSize: "1rem",
                  fontWeight: 600,
                  "&:hover": {
                    borderColor: "#303f9f",
                    borderWidth: 2,
                    backgroundColor: "rgba(63, 81, 181, 0.04)",
                  },
                }}
              >
                ログイン
              </Button>
            </Box>
          </Box>
        </Paper>
      </Container>
    </div>
  );
};

export default Signup;
