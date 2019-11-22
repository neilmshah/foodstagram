import imagesStyles from "assets/jss/material-kit-react/imagesStyles.js";
import { createMuiTheme } from "@material-ui/core/styles";
import blue from "@material-ui/core/colors/blue";

const theme = createMuiTheme({
  palette: {
    primary: blue
  }
});

const timelinePageStyle = {
  root: {
    flexGrow: 1
  },
  modal: {
    display: "grid",
    alignItems: "center",
    justifyContent: "center",
    // display: "grid"
    overflow: "scroll"
  },
  paper: {
    backgroundColor: theme.palette.background.paper,
    border: "2px solid #000",
    boxShadow: theme.shadows[5],
    padding: theme.spacing(2, 4, 3)
  },
  comment: {
    overflow: "scroll",
    height: "700px",
    padding: "2rem",
    paddingBottom: "10rem"
  },
  header: {
    textAlign: "center"
  },
  comments_title: {
    paddingTop: "1rem",
    textAlign: "center",
    fontWeight: "bolder"
  },
  description: {
    overflow: "scroll",
    // height: "400px",
    padding: "2rem"
  },
  ...imagesStyles,
  textMuted: {
    color: "#6c757d"
  },
  media: {
    height: 0,
    paddingTop: "56.25%" // 16:9
  },
  card: {
    maxWidth: 600
  },
  marginLeft: {
    left: 0
  },
  marginRight: {
    float: "right"
  }
};

export default timelinePageStyle;
