import React from "react";
import axios from "axios";
// material-ui components
import { withStyles } from "@material-ui/core/styles";
// core components
import Badge from "@material-ui/core/Badge";
import Card from "components/Card/Card.js";
import CardBody from "components/Card/CardBody.js";
// core components
import Button from "components/CustomButtons/Button.js";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";
import Fab from "@material-ui/core/Fab";
import AddCommentIcon from "@material-ui/icons/AddCommentOutlined";
import FavoriteIcon from "@material-ui/icons/Favorite";
import Header from "components/Header/Header.js";
import HeaderLinks from "components/Header/HeaderLinks.js";
import timelinePageStyle from "assets/jss/material-kit-react/views/timelinePage";
import Grid from "@material-ui/core/Grid";
import Modal from "@material-ui/core/Modal";
import Fade from "@material-ui/core/Fade";
import Backdrop from "@material-ui/core/Backdrop";
import Box from "@material-ui/core/Box";
import TextField from "@material-ui/core/TextField";
import * as CONST from "../../config";
import PropTypes from "prop-types";
import Moment from "react-moment";
class TimeLinePage extends React.Component {
  static propTypes = {
    history: PropTypes.object.isRequired
  };

  constructor(props) {
    super(props);
    this.state = {
      timeline: [],
      comments: [],
      open: false,
      showDialog: false,
      userid: localStorage.getItem("user_id"),
      fullname: localStorage.getItem("user_fullname"),
      hasLiked: false,
      selectedPost: {},
      commentMsg: ""
    };
    this.handlesOnPost = this.handlesOnPost.bind(this);
    this.handleOpen = this.handleOpen.bind(this);
    this.handlesOnCommentChange = this.handlesOnCommentChange.bind(this);
    this.handleLike = this.handleLike.bind(this);
  }
  componentDidMount() {
    this.getTimeline();
  }

  getTimeline = () => {
    axios
      .get(CONST.TIMELINE_SERVICE + "/timeline", {
        headers: {
          "x-api-key": process.env.REACT_APP_TIMELINE_SERVICE_AUTH_KEY
        }
      })
      .then(res => {
        console.log("Status: " + res.status);
        console.log("Data: " + JSON.stringify(res.data));
        var timeline = res.data.Timeline;
        if (timeline.length != 0)
          timeline.sort(
            (a, b) => parseInt(b.Timestamp) - parseInt(a.Timestamp)
          );
        if (res.status == 200) {
          this.setState({
            timeline: timeline
          });
        } else {
          this.getAllPosts();
        }
      })
      .catch(e => {
        this.getAllPosts();
      });
  };

  getAllPosts = () => {
    axios
      .get(CONST.IMAGE_SERVICE + "/image", {
        headers: {
          apikey: process.env.REACT_APP_IMAGE_SERVICE_AUTH_KEY
        }
      })
      .then(res => {
        console.log("Status: " + res.status);
        console.log("Data: " + JSON.stringify(res.data));
        var timeline = res.data;
        if (timeline.length != 0)
          timeline.sort(
            (a, b) => parseInt(b.Timestamp) - parseInt(a.Timestamp)
          );
        if (res.status == 200) {
          this.setState({
            timeline: timeline
          });
        }
      });
  };

  getCommentsForImage = imageId => {
    axios
      .get(
        CONST.COMMENTS_LIKES_SERVICE +
          "/comment/" +
          imageId +
          "/" +
          this.state.userid
      )
      .then(res => {
        if (res.status == 200) {
          var commentList = res.data.Comments;
          if (commentList.length != 0)
            commentList.sort(
              (a, b) => parseInt(a.Timestamp) - parseInt(b.Timestamp)
            );
          this.setState({
            comments: commentList,
            hasLiked: res.data.Liked
          });
        }
      });
  };

  handleLike = () => {
    if (this.state.hasLiked) return;
    var body = {
      User_id: this.state.userid
    };
    axios
      .post(
        CONST.COMMENTS_LIKES_SERVICE + "/like/" + this.state.selectedPost.Id,
        body
      )
      .then(res => {
        if (res.status == 200) {
          this.setState({ hasLiked: true });
        }
      });
  };
  handlesOnPost = () => {
    if (this.state.commentMsg == "") return;
    var myComment = this.state.commentMsg;
    var body = {
      User_id: this.state.userid,
      User_name: this.state.fullname,
      Comment: myComment
    };
    axios
      .post(
        CONST.COMMENTS_LIKES_SERVICE + "/comment/" + this.state.selectedPost.Id,
        body
      )
      .then(res => {
        console.log("Status: " + res.status);
        console.log("Data: " + JSON.stringify(res.data));
        if (res.status == 200) {
          var commentObj = {
            User_name: this.state.fullname,
            Comment: myComment,
            Timestamp: Date.parse(new Date().toISOString().valueOf()) / 1000
          };
          var comments = this.state.comments;
          comments.push(commentObj);
          this.setState({ comments: comments, commentMsg: "" });
        }
      });
  };

  handleOpen = post => {
    if (!this.state.userid || this.state.userid == "") {
      this.setState({ showDialog: true });
    } else {
      this.getCommentsForImage(post.Id);
      console.log("Image ID::", post.Id);
      this.setState({ open: true, selectedPost: post });
    }
  };

  handleClose = () => {
    this.setState({
      open: false,
      selectedPost: {},
      comments: [],
      commentMsg: ""
    });
    this.getTimeline();
  };

  handleDialogClose = () => {
    this.props.history.push("/login");
  };

  handlesOnCommentChange = (event: SyntheticEvent<>) => {
    if (event) {
      var commentMsg = event.target.value;
      this.setState({ commentMsg: commentMsg });
    }
  };

  render() {
    const { classes, ...rest } = this.props;
    const { open } = this.state;
    return (
      <div>
        <Header
          color="info"
          brand="Foodstagram"
          rightLinks={<HeaderLinks />}
          fixed
          {...rest}
        />
        <Box m={10} />
        <Grid
          container
          direction="column"
          alignItems="center"
          justify="center"
          style={{ minHeight: "100vh" }}
        >
          {this.state.timeline.map((post, index) => {
            return (
              <Grid
                item
                xs={6}
                spacing={0}
                onClick={() => this.handleOpen(post)}
              >
                <Card className={classes.card}>
                  <img
                    className={classes.imgCardTop}
                    src={post.Url}
                    alt="Card-img-cap"
                  />
                  <CardBody>
                    <h3 className={classes.cardTitle}>{post.UserName}</h3>
                    <p>{post.Description}</p>
                    <p>
                      <small className={classes.textMuted}>
                        <Moment fromNow unix>
                          {post.Timestamp}
                        </Moment>
                      </small>
                    </p>
                    <Badge
                      className={classes.marginLeft}
                      badgeContent={post.LikeCount}
                      color="primary"
                    >
                      <FavoriteIcon />
                    </Badge>
                    <Badge
                      className={classes.marginRight}
                      badgeContent={post.CommentCount}
                      color="primary"
                    >
                      <AddCommentIcon />
                    </Badge>
                  </CardBody>
                </Card>
              </Grid>
            );
          })}
          <Modal
            aria-labelledby="transition-modal-title"
            aria-describedby="transition-modal-description"
            className={classes.modal}
            open={open}
            onClose={this.handleClose}
            closeAfterTransition
            BackdropComponent={Backdrop}
            BackdropProps={{
              timeout: 500
            }}
          >
            <Fade in={open}>
              <div className={classes.paper}>
                <h2 id="transition-modal-title" className={classes.header}>
                  {this.state.selectedPost.UserName}
                </h2>
                <Grid container className={classes.root}>
                  <Grid item xs={6}>
                    <Card className={classes.card}>
                      <img
                        className={classes.imgCardTop}
                        src={this.state.selectedPost.Url}
                        alt="Card-img-cap"
                      />
                      <CardBody className={classes.description}>
                        <p>{this.state.selectedPost.Description}</p>
                      </CardBody>
                    </Card>
                    <Fab
                      aria-label="like"
                      className={classes.fab}
                      color="default"
                      onClick={this.handleLike}
                    >
                      <FavoriteIcon
                        style={{
                          color: this.state.hasLiked ? "red" : "black"
                        }}
                        className="material-icons"
                      />
                    </Fab>
                  </Grid>
                  <Grid item xs={6}>
                    <h4 className={classes.comments_title}>Comments</h4>
                    <div className={classes.comment}>
                      {this.state.comments.map((comment, index) => {
                        return (
                          <Box>
                            <h6>{comment.User_name}</h6>
                            <p>{comment.Comment}</p>
                            <p>
                              <small className={classes.textMuted}>
                                <Moment fromNow unix>
                                  {comment.Timestamp}
                                </Moment>
                              </small>
                            </p>
                          </Box>
                        );
                      })}
                      <TextField
                        id="commentMsg"
                        label={this.state.fullname}
                        multiline
                        fullWidth
                        rows="4"
                        placeholder="Your comments"
                        value={this.state.commentMsg}
                        className={classes.textField}
                        margin="normal"
                        variant="outlined"
                        InputLabelProps={{
                          shrink: true
                        }}
                        inputProps={{
                          onChange: this.handlesOnCommentChange,
                          type: "commentMsg"
                        }}
                      />
                      <Button
                        size="lg"
                        onClick={() => {
                          console.log("calling handlesOnPost");
                          this.handlesOnPost();
                        }}
                      >
                        Post
                      </Button>
                    </div>
                  </Grid>
                </Grid>
              </div>
            </Fade>
          </Modal>
        </Grid>

        <Dialog
          open={this.state.showDialog}
          onClose={this.handleDialogClose}
          aria-labelledby="responsive-dialog-title"
        >
          <DialogContent>
            <DialogContentText>
              Please sign in to continue exploring Foodstagram.
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleDialogClose} color="info" autoFocus>
              OK
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default withStyles(timelinePageStyle)(TimeLinePage);
