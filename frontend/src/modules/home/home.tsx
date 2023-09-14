import { PrimaryButton, Image, Text } from "@fluentui/react";
import { useNavigate } from "react-router-dom";
import { initializeIcons } from "@fluentui/font-icons-mdl2";
import { Icon } from "@fluentui/react/lib/Icon";
import { Wrapper, Row, LeftColumn, RightColumn } from "./home.styles";
import Footer from "../../components/footer";
const Home = () => {
  initializeIcons();

  const navigate = useNavigate();
  const start = async () => {
    // window.location.href ="https://project-zen.azurewebsites.net/auth/start-oauth-flow";
    navigate("/dashboard");
    /*const response = await fetch(
      "https://project-zen.azurewebsites.net/auth/start-oauth-flow"
    );
    if (
      (response as any).access_token ||
      (response as any).message == "Authorization successful"
    ) {
      navigate("/dashboard");
    } else {
      window.location.href =
        "https://project-zen.azurewebsites.net/auth/start-oauth-flow";
    }*/
  };

  return (
    <Wrapper>
      <Row>
        <LeftColumn>
          <Image
            src="logo.png" // Add your logo image here
            style={{ width: 100, margin: "auto", marginBottom: "50px" }}
          />
          <Text
            as="h1"
            style={{
              fontSize: 38,
            }}
          >
            Welcome to <b>Project-Zen</b>
          </Text>
          <br />
          <Text
            as="h3"
            style={{
              fontSize: 28,
              marginBottom: "40px",
              marginTop: 15,
              color: "#666",
            }}
          >
            Uncluttered Inboxes mean Happy Souls
          </Text>

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
          <PrimaryButton
            primary
            style={{ width: 150, height: 50, fontSize: 16 }}
            href="https://project-zen.azurewebsites.net/auth/start-oauth-flow"
          >
            Get Started
          </PrimaryButton>
        </LeftColumn>
        <RightColumn>
          <Image src="home.png" style={{ width: "100%" }} />
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
