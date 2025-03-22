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
  useTheme,
  Grid,
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
import { useAuth } from "../contexts/AuthContext";

const Signup: React.FC = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const { signup } = useAuth();
  const navigate = useNavigate();
  const theme = useTheme();

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleClickShowConfirmPassword = () => {
    setShowConfirmPassword(!showConfirmPassword);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    // バリデーション
    if (!name.trim()) {
      return setError("お名前を入力してください");
    }

    if (password !== confirmPassword) {
      return setError("パスワードが一致しません");
    }

    if (password.length < 6) {
      return setError("パスワードは6文字以上にしてください");
    }

    setLoading(true);

    try {
      await signup(name, email, password);
      navigate("/home");
    } catch (error: any) {
      setError(error.message || "ユーザー登録に失敗しました");
    } finally {
      setLoading(false);
    }
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

            <form onSubmit={handleSubmit}>
              <TextField
                label="お名前"
                fullWidth
                margin="normal"
                value={name}
                onChange={(e) => setName(e.target.value)}
                autoComplete="name"
                autoFocus
              />

              <TextField
                label="メールアドレス"
                type="email"
                fullWidth
                margin="normal"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                autoComplete="email"
              />

              <TextField
                label="パスワード"
                type={showPassword ? "text" : "password"}
                fullWidth
                margin="normal"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                autoComplete="new-password"
                InputProps={{
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
              />

              <TextField
                label="パスワード確認"
                type={showConfirmPassword ? "text" : "password"}
                fullWidth
                margin="normal"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                autoComplete="new-password"
                InputProps={{
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
                disabled={loading}
              >
                {loading ? "登録中..." : "アカウント作成"}
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
