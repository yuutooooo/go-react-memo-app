import React, { useState } from "react";
import {
  Typography,
  Container,
  Box,
  Paper,
  TextField,
  Button,
  Alert,
  InputAdornment,
  IconButton,
  Divider,
  useMediaQuery,
  useTheme,
  Grid,
} from "@mui/material";
import {
  Email as EmailIcon,
  Lock as LockIcon,
  Visibility as VisibilityIcon,
  VisibilityOff as VisibilityOffIcon,
  ArrowBack as ArrowBackIcon,
} from "@mui/icons-material";
import { useNavigate, Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

const Login: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const { login } = useAuth();
  const navigate = useNavigate();
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down("sm"));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      await login(email, password);
      navigate("/home");
    } catch (error: any) {
      setError(error.message || "ログインに失敗しました");
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
              アカウントにログイン
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

            <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
              <TextField
                margin="normal"
                required
                fullWidth
                label="メールアドレス"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                autoComplete="email"
                autoFocus
              />
              
              <TextField
                margin="normal"
                required
                fullWidth
                label="パスワード"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                autoComplete="current-password"
              />
              
              <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2, py: 1.5 }}
                disabled={loading}
              >
                {loading ? 'ログイン中...' : 'ログイン'}
              </Button>
            </Box>

            <Divider sx={{ my: 3 }}>
              <Typography variant="body2" color="textSecondary">
                または
              </Typography>
            </Divider>

            <Box className="flex justify-center">
              <Button
                component={Link}
                to="/signup"
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
                新規アカウント登録
              </Button>
            </Box>

            <Box className="mt-4 text-center">
              <Typography
                variant="body2"
                sx={{ color: "text.secondary", fontStyle: "italic", mt: 2 }}
              >
                テスト用アカウント: test@example.com / password
              </Typography>
            </Box>
          </Box>
        </Paper>
      </Container>
    </div>
  );
};

export default Login;
