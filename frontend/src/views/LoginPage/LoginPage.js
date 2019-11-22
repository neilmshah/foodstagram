import React from "react";
import axios from "axios"
// @material-ui/core components
import InputAdornment from "@material-ui/core/InputAdornment";
import Icon from "@material-ui/core/Icon";
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
import PropTypes from "prop-types";
import loginPageStyle from "assets/jss/material-kit-react/views/loginPage.js";
import withStyles from "@material-ui/core/styles/withStyles";
import image from "assets/img/food1.jpeg";
import { stringify } from "querystring";
import * as CONST from "../../config";


class LoginPage extends React.Component {
  
  static propTypes = {
    history: PropTypes.object.isRequired
  };

  constructor(props) {
    super(props);
    // we use this to make the card to appear after the page has been rendered
    this.state = {
      username: "",
      password: "",
      redirectToHome:false
    };
  }
  
  render() {
    if(this.state.redirectToHome)
      this.props.history.push("/home");
    const {classes, ...rest } = this.props;
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
                      <h4>Login</h4>
                    </CardHeader>
                    <CardBody>
                      
                      <CustomInput
                        labelText="Username"
                        id="username"
                        formControlProps={{
                          fullWidth: true
                        }}
                        inputProps={{
                          onChange: this.handlesOnUsernameChange,
                          type: "text",
                          endAdornment: (
                            <InputAdornment position="end">
                              <People className={classes.inputIconsColor} />
                            </InputAdornment>
                          )
                        }}
                      />
                      <CustomInput
                        labelText="Password"
                        id="password"
                        formControlProps={{
                          fullWidth: true
                        }}
                        inputProps={{
                          onChange: this.handlesOnPasswordChange,
                          type: "password",
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
                      <Button simple color="success" size="lg" onClick={this.handlesOnLogin}>
                        Get started
                      </Button>
                    </CardFooter>

                    <CardFooter className={classes.cardFooter}>
                    New to Foodstagram?
                    <a
                      href=""
                      onClick={e => {
                        this.props.history.push("/signup");
                      }}
                    >
                    {"   "}Join now
                    </a>
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
  handlesOnLogin = (event: SyntheticEvent<>) => {
    const { history } = this.props;
    const { username, password } = this.state;
    var body = {
      username: username,
      password: password
    }

    axios.post(CONST.USER_PROFIE_SERVICE+'/login', body)
    .then(res=> {
      if(res.status !=200){
        alert("Service Error");
      }else{
        if(res.data.error) {
          alert(res.data.error)
          return
        }
        var user = res.data
        if (user== null || user.username == null){
          alert("Username cannot be null");
          return
        }
        localStorage.username = user.username;
        localStorage.user_fullname = user.firstname + " " + user.lastname
        localStorage.user_id = user.userID;
        localStorage.token = user.token
        this.setState({redirectToHome:true})
      }
     
    })
    .catch(function (error) {
        alert("Error :" +stringify(error))
    });

  };

  handlesOnPasswordChange = (event: SyntheticEvent<>) => {
    if (event) {
      this.setState({ password: event.target.value.trim() });
    }
  };
  

  handlesOnUsernameChange = (event: SyntheticEvent<>) => {
    if (event) {
      this.setState({ username: event.target.value.trim() });
    }
  };
}

export default withStyles(loginPageStyle)(LoginPage);
