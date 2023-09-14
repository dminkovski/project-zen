import { Image } from "@fluentui/react";
import Footer from "../../components/footer";

const About = () => {
  return (
    <div style={{ textAlign: "center", padding: "50px" }}>
      <Image
        src="logo.png" // Add your logo image here
        style={{ width: 100, margin: "auto", marginBottom: "50px" }}
      />
      <h1>About Us</h1>
      <p>
        We are a dynamic team of four passionate individuals brought together by
        a common goal: to revolutionize the way people manage their inboxes.
      </p>
      <p>
        Our diverse skills and backgrounds, ranging from software development to
        sarcasm and meme creation, have allowed us to create Project-Zen - an
        automated tool designed to declutter and simplify your inbox experience.
      </p>
      <p>
        With Project-Zen, we aim to bring a sense of calm and efficiency to your
        digital life, ensuring you never miss out on important emails while
        keeping the noise at bay.
      </p>
      <p>Best regards from Sumit, Karim, Hanna & David</p>
      <Footer />
    </div>
  );
};

export default About;
