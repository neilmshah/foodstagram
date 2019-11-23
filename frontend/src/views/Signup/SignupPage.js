import React from "react";
import axios from "axios";
import * as CONST from "../../config";
// @material-ui/core components
import InputAdornment from "@material-ui/core/InputAdornment";
import Icon from "@material-ui/core/Icon";
import withStyles from "@material-ui/core/styles/withStyles";
// @material-ui/icons

import People from "@material-ui/icons/People";
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
import CustomInput from "components/CustomInput/CustomInput.js";
import signupPageStyle from "assets/jss/material-kit-react/views/loginPage.js";
import PropTypes from "prop-types";
import image from "assets/img/food1.jpeg";
import { stringify } from "querystring";

class SignupPage extends React.Component {
  static propTypes = {
    history: PropTypes.object.isRequired
  };

  constructor(props) {
    super(props);
    // we use this to make the card to appear after the page has been rendered
    this.state = {
      firstname: "",
      lastname: "",
      username: "",
      password: "",
      redirectToLogin: false
    };
  }

  render() {
    if (this.state.redirectToLogin) this.props.history.push("/login");
    const { classes, ...rest } = this.props;
    return (
      <div>
        <Header
          absolute
          color="transparent"
          brand="Foodstagram"
          rightLinks={<HeaderLinks />}
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
              <GridItem xs={12} sm={12} md={4}>
                <Card>
                  <form className={classes.form}>
                    <CardHeader color="success" className={classes.cardHeader}>
                      <h4>Enter details</h4>
                    </CardHeader>

                    <CardBody>
                      <CustomInput
                        labelText="First Name..."
                        id="first"
                        formControlProps={{
                          fullWidth: true
                        }}
                        inputProps={{
                          type: "text",
                          onChange: this.handlesOnFirstnameChange,
                          endAdornment: (
                            <InputAdornment position="end">
                              <People className={classes.inputIconsColor} />
                            </InputAdornment>
                          )
                        }}
                      />
                      <CustomInput
                        labelText="Last Name..."
                        id="first"
                        formControlProps={{
                          fullWidth: true
                        }}
                        inputProps={{
                          type: "text",
                          onChange: this.handlesOnLastnameChange,
                          endAdornment: (
                            <InputAdornment position="end">
                              <People className={classes.inputIconsColor} />
                            </InputAdornment>
                          )
                        }}
                      />
                      <CustomInput
                        labelText="Username..."
                        id="first"
                        formControlProps={{
                          fullWidth: true
                        }}
                        inputProps={{
                          type: "text",
                          onChange: this.handlesOnUsernameChange,
                          endAdornment: (
                            <InputAdornment position="end">
                              <People className={classes.inputIconsColor} />
                            </InputAdornment>
                          )
                        }}
                      />
                      <CustomInput
                        labelText="Password"
                        id="pass"
                        formControlProps={{
                          fullWidth: true
                        }}
                        inputProps={{
                          type: "password",
                          onChange: this.handlesOnPasswordChange,
                          endAdornment: (
                            <InputAdornment position="end">
                              <Icon className={classes.inputIconsColor}>
                                lock_outline
                              </Icon>
                            </InputAdornment>
                          ),
                          autoComplete: "off"
                        }}
                      />
                    </CardBody>
                    <CardFooter className={classes.cardFooter}>
                      <Button
                        simple
                        color="success"
                        size="lg"
                        onClick={this.handlesOnSignup}
                      >
                        Sign Up
                      </Button>
                    </CardFooter>
                  </form>
                </Card>
              </GridItem>
            </GridContainer>
          </div>
          <Footer whiteFont />
        </div>
      </div>
    );
  }

  handlesOnPasswordChange = (event: SyntheticEvent<>) => {
    if (event) {
      this.setState({ password: event.target.value.trim() });
    }
  };

  handlesOnSignup = (event: SyntheticEvent<>) => {
    const { history } = this.props;
    const { firstname, lastname, username, password } = this.state;
    var body = {
      username: username,
      password: password,
      firstname: firstname,
      lastname: lastname
    };

    axios
      .post(CONST.USER_PROFIE_SERVICE + "/register", body)
      .then(res => {
        if (res.status != 200) {
          alert("Service Error");
        } else {
          var resObj = res.data;
          if (resObj.result != "Registration Successful") {
            alert("Username already exist");
            return;
          }

          this.setState({ redirectToLogin: true });
        }
      })
      .catch(function(error) {
        alert("Error occured. Please try again!");
      });
  };

  handlesOnUsernameChange = (event: SyntheticEvent<>) => {
    if (event) {
      this.setState({ username: event.target.value.trim() });
    }
  };

  handlesOnFirstnameChange = (event: SyntheticEvent<>) => {
    if (event) {
      this.setState({ firstname: event.target.value.trim() });
    }
  };
  handlesOnLastnameChange = (event: SyntheticEvent<>) => {
    if (event) {
      this.setState({ lastname: event.target.value.trim() });
    }
  };
}
export default withStyles(signupPageStyle)(SignupPage);
