/*eslint-disable*/
import React from "react";
import DeleteIcon from "@material-ui/icons/Delete";
import IconButton from "@material-ui/core/IconButton";
// react components for routing our app without refresh
import { Link } from "react-router-dom";

// @material-ui/core components
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Tooltip from "@material-ui/core/Tooltip";

// @material-ui/icons
import { Apps, AccountCircle, Image } from "@material-ui/icons";

// core components
import CustomDropdown from "components/CustomDropdown/CustomDropdown.js";
import Button from "components/CustomButtons/Button.js";

import styles from "assets/jss/material-kit-react/components/headerLinksStyle.js";

const useStyles = makeStyles(styles);

export default function HeaderLinks(props) {
  const classes = useStyles();

  const clearLocalStorage = () => {
   localStorage.clear()
};
  return (
    <List className={classes.list}>

      <ListItem className={classes.listItem}>
        <Button
          href="/image"
          color="transparent"
          className={classes.navLink}
        >
          <Image className={classes.icons} /> Post
        </Button>
      </ListItem>
      <ListItem className={classes.listItem}>
       {localStorage.getItem('user_id') && 
          <Button
            href="/profile"
            color="transparent"
            className={classes.navLink}
          >
            <AccountCircle className={classes.icons} /> {localStorage.getItem('username')}
          </Button>
        }
        {!localStorage.getItem('user_id') && 
          <Button
            href="/profile"
            color="transparent"
            className={classes.navLink}
          >
            <AccountCircle className={classes.icons} /> Profile
          </Button>
        }
      </ListItem>
      { localStorage.getItem("user_id") &&       
      <ListItem className={classes.listItem}>
        <Tooltip
          id="instagram-tooltip"
          title="Sign Out"
          placement={window.innerWidth > 959 ? "top" : "left"}
          classes={{ tooltip: classes.tooltip }}
        >
          <Button
            color="transparent"
            href="/login"
            className={classes.navLink}
            onClick={clearLocalStorage}
          >
            Sign Out
            {/* <i className={classes.socialIcons + " fab fa-instagram"} /> */}
          </Button>
        </Tooltip>
      </ListItem>}


      {!localStorage.getItem("user_id") && <ListItem className={classes.listItem}>
        <Tooltip
          id="instagram-tooltip"
          title="Sign In"
          placement={window.innerWidth > 959 ? "top" : "left"}
          classes={{ tooltip: classes.tooltip }}
        >
          <Button
            color="transparent"
            taget= "_self"
            href="/login"
            className={classes.navLink}
          >
            Sign In
          </Button>
        </Tooltip>
      </ListItem>}
      
    </List>
  );
}