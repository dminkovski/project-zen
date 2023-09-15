import { Wrapper, Row, LeftColumn, RightColumn } from "./home.styles";
import Footer from "../../components/Footer/Footer";
import GoogleButton from "../../components/GoogleButton/GoogleButton";
import Logo from "../../components/Logo/Logo";
import * as Config from "../../config/config";
const Home = () => {
  return (
    <Wrapper>
      <Row>
        <LeftColumn>
          <Logo />
          <h1
            style={{
              fontSize: 38,
            }}
          >
            Welcome to <b>Project-Zen</b>
          </h1>
          <br />
          <h3
            style={{
              fontSize: 28,
              marginBottom: "40px",
              marginTop: 15,
              color: "#666",
            }}
          >
            Uncluttered Inboxes mean Happy Souls
          </h3>

          <p
            style={{
              color: "#666",
              marginTop: "50px",
              marginBottom: "30px",
              fontSize: "18px",
              textAlign: "center",
            }}
          >
            Is your inbox resembling a Black Friday sale flyer after a
            hurricane?
            <br />
            <br />
            Fear not! Our <b style={{ color: "#0078d4" }}>AI Robot</b> is here
            to save the day!
            <br />
            Say goodbye to clutter, and hello to peace of mind!
          </p>
          <GoogleButton href={Config.AUTHENTICATION_URL}>
            Continue with Google
          </GoogleButton>
        </LeftColumn>
        <RightColumn>
          <img src="home.png" style={{ width: "100%" }} />
          <span style={{ color: "#666", fontStyle: "italic" }}>
            Emails opened by a cute robot standing in front of a zen garden,
            digital art
          </span>
        </RightColumn>
      </Row>
      <Footer />
    </Wrapper>
  );
};
export default Home;
