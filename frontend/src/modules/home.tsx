import { AutoScroll, Button, Image, Text } from "@fluentui/react";

const Home = () => {
  return (
    <div>
      <div>
        <Text
          as="h1"
          style={{
            fontSize: 38,
          }}
        >
          Welcome to Project-Zen
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
        <Image
          src="hero.png" // Add your logo image here
          style={{ marginBottom: "30px", width: 300, margin: "auto" }}
        />
        <p
          style={{
            color: "#666",
            width: 800,
            margin: "auto",
            marginBottom: "30px",
            fontSize: "18px",
            textAlign: "center",
          }}
        >
          Is your inbox resembling a Black Friday sale flyer after a hurricane?
          <br />
          <br />
          Fear not! Project-Zen is here to save the day!
          <br />
          Say goodbye to clutter, and hello to serenity!
        </p>
        <h2 style={{ marginTop: 50 }}>Why Zen is great?</h2>
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            marginTop: "30px",
            alignContent: "center",
          }}
        >
          <div>
            <h3>Tame the Email Tornado</h3>
            <p>
              Never get lost in your inbox again.
              <br />
              We bring order to the chaos and get rid of newsletters!
            </p>
          </div>
          <div>
            <h3>Automation at Its Best</h3>
            <p>
              Let the magic happen without lifting your finger. <br />
              Our software does everything for you.
            </p>
          </div>
          <div>
            <h3>Save Big with Ease</h3>
            <p>
              Discounts and gifts are just a click away
              <br />, waiting for you to claim them.
            </p>
          </div>
          <div>
            <h3>Stay Informed, Stress-Free</h3>
            <p>
              Your customized dashboard ensures you're always in control. <br />
              Say goodbye to FOMO and hello to peace of mind!
            </p>
          </div>
        </div>
        <Button style={{ marginTop: 30, marginBottom: 30 }} primary>
          Get Started
        </Button>
        <p style={{ marginTop: "20px", color: "#666", fontSize: "14px" }}>
          Already using Project-Zen? <a href="/">Leave feedback</a>
        </p>
        <br />
        <div className="ui container">
          <p>
            <a href="/about" style={{ color: "#666", marginRight: 5 }}>
              | About Us |
            </a>
            <a href="/about" style={{ marginLeft: 5, color: "#666" }}>
              Legal |
            </a>
          </p>
          <p> &copy; 2023 Project-Zen. All rights reserved.</p>
        </div>
      </div>
    </div>
  );
};
export default Home;