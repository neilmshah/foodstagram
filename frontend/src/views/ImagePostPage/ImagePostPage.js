import React from "react";
import Icon from "@material-ui/core/Icon";
// core components
import Header from "components/Header/Header.js";
import HeaderLinks from "components/Header/HeaderLinks.js";
import Footer from "components/Footer/Footer.js";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem.js";
import Button from "components/CustomButtons/Button.js";
import Card from "components/Card/Card.js";
import CardBody from "components/Card/CardBody.js";
import CardHeader from "components/Card/CardHeader.js";
import CardFooter from "components/Card/CardFooter.js";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";
import TextField from "@material-ui/core/TextField";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import PropTypes from "prop-types";
import loginPageStyle from "assets/jss/material-kit-react/views/loginPage.js";
import withStyles from "@material-ui/core/styles/withStyles";
import image from "assets/img/food1.jpeg";
import * as CONST from "../../config";

class ImagePostPage extends React.Component {
  static propTypes = {
    history: PropTypes.object.isRequired
  };

  constructor(props) {
    super(props);
    // we use this to make the card to appear after the page has been rendered
    this.state = {
      username: localStorage.username,
      userid: localStorage.user_id,
      fullname: localStorage.getItem("user_fullname"),
      foodImage: null,
      description: "",
      imageSrc: null,
      responseSuccessText: null,
      responseErrorText: null,
      showDialog: false
    };
  }

  handleDialogClose = () => {
    if (this.state.responseErrorText == null) this.props.history.push("/home");
    this.setState({ showDialog: false });
  };

  render() {
    const { classes, ...rest } = this.props;
    return (
      <div>
        <Header
          color="success"
          brand="Foodstagram"
          rightLinks={<HeaderLinks />}
          fixed
          {...rest}
        />
        <div
          className={classes.pageHeader}
          style={{
            backgroundImage: "url(" + image + ")",
            backgroundSize: "cover",
            backgroundPosition: "top center"
          }}
        >
          <div className={classes.container}>
            <GridContainer justify="center">
              <GridItem xs={12} sm={12} md={8}>
                <Card>
                  <form className={classes.form}>
                    <CardHeader color="success" className={classes.cardHeader}>
                      <h4>Post Feed</h4>
                    </CardHeader>
                    <CardBody>
                      <TextField
                        id="description"
                        label="Description"
                        multiline
                        fullWidth
                        rows="10"
                        placeholder="Image Description"
                        className={classes.textField}
                        margin="normal"
                        variant="outlined"
                        inputProps={{
                          onChange: this.handlesOnDescriptionChange,
                          type: "description"
                        }}
                        InputLabelProps={{
                          shrink: true
                        }}
                      />

                      <Box
                        width="100%"
                        marginBottom="20px"
                        style={{
                          display: this.state.imageSrc != null,
                          justifyContent: "center"
                        }}
                      >
                        <img src={this.state.imageSrc} width="100%" />
                        {this.state.foodImage ? this.state.foodImage.name : ""}
                      </Box>

                      <input
                        accept="image/*"
                        type="file"
                        onChange={this.handlesOnImageChange("image")}
                        style={{ display: "none" }}
                        id="icon-button-file"
                      />
                      <label htmlFor="icon-button-file">
                        <Button
                          variant="raised"
                          color="github"
                          component="span"
                        >
                          Upload
                          <Icon className={this.props.classes.rightIcon}>
                            file_upload
                          </Icon>
                        </Button>
                      </label>
                    </CardBody>
                    <CardFooter className={classes.cardFooter}>
                      <Button
                        variant="raised"
                        color="success"
                        component="span"
                        // size="lg"
                        onClick={this.handlesOnPost}
                      >
                        Post
                      </Button>
                    </CardFooter>
                  </form>
                </Card>
              </GridItem>
            </GridContainer>
          </div>
          <Footer whiteFont />
        </div>
        <Dialog
          open={this.state.showDialog}
          onClose={this.handleDialogClose}
          aria-labelledby="responsive-dialog-title"
        >
          <DialogContent>
            <DialogContentText>
              {this.state.responseSuccessText == null
                ? this.state.responseErrorText
                : this.state.responseSuccessText}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button
              onClick={this.handleDialogClose}
              color={
                this.state.responseSuccessText == null ? "danger" : "success"
              }
              autoFocus
            >
              OK
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }

  handlesOnPost = (event: SyntheticEvent<>) => {
    console.log(
      "process.env.REACT_APP_IMAGE_SERVICE_AUTH_KEY::",
      process.env.REACT_APP_IMAGE_SERVICE_AUTH_KEY
    );
    const { fullname, userid, foodImage, description } = this.state;
    var data = new FormData();
    data.set("username", fullname);
    data.set("userid", userid);
    data.append("foodImage", foodImage);
    data.set("description", description);
    fetch(CONST.IMAGE_SERVICE + "/image", {
      method: "POST",
      headers: new Headers({
        apikey: process.env.REACT_APP_IMAGE_SERVICE_AUTH_KEY
      }),
      body: data
    })
      .then(response => {
        if (response.status == 200) {
          this.setState({
            // showPostButton: false,
            showDialog: true,
            responseSuccessText: "Image uploaded successfully!!",
            responseErrorText: null
          });
        } else {
          this.setState({
            showDialog: true,
            responseSuccessText: null,
            responseErrorText:
              "Error posting the data. Please try again after some time!"
          });
        }
        // return response.json();
      })
      .catch(e => {
        console.log(e);
        this.setState({
          showDialog: true,
          responseSuccessText: null,
          responseErrorText:
            "Error posting the data. Please try again after some time!"
        });
      });
  };

  handlesOnImageChange = name => event => {
    const value = name === "image" ? event.target.files[0] : event.target.value;
    this.setState({ ["foodImage"]: value });
    this.setState({ imageSrc: URL.createObjectURL(value) });
  };

  handlesOnDescriptionChange = (event: SyntheticEvent<>) => {
    if (event) {
      var description = event.target.value.trim();
      this.setState({ description: description });
    }
  };
}

export default withStyles(loginPageStyle)(ImagePostPage);
