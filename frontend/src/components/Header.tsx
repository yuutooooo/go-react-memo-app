import { AppBar, Toolbar, Typography, IconButton, Box, Badge, Tooltip } from "@mui/material"
import { FolderOutlined, DescriptionOutlined, NotificationsOutlined } from "@mui/icons-material"
import LogoutButton from "./LogoutButton"

const Header = () => {
  return (
    <AppBar position="static" color="default" elevation={1}>
      <Toolbar>
        <Typography 
          variant="h6" 
          component="div" 
          sx={{ 
            flexGrow: 1,
            fontWeight: 'bold',
            color: '#3f51b5'
          }}
        >
          メモアプリ
        </Typography>
        
        <Box sx={{ display: 'flex', gap: 1, mr: 2 }}>
          <Tooltip title="フォルダ">
            <IconButton color="primary">
              <Badge badgeContent={4} color="primary">
                <FolderOutlined />
              </Badge>
            </IconButton>
          </Tooltip>
          
          <Tooltip title="ノート">
            <IconButton color="primary">
              <Badge badgeContent={3} color="error">
                <DescriptionOutlined />
              </Badge>
            </IconButton>
          </Tooltip>
          
          <Tooltip title="通知">
            <IconButton color="primary">
              <NotificationsOutlined />
            </IconButton>
          </Tooltip>
        </Box>
        
        <LogoutButton variant="outlined" color="primary" />
      </Toolbar>
    </AppBar>
  )
};

export default Header;
