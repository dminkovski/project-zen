import Footer from "../../components/Footer/Footer";
import Logo from "../../components/Logo/Logo";

const Legal = () => {
  return (
    <div style={{ textAlign: "center", padding: "50px" }}>
      <Logo />
      <h1>Legal</h1>
      <p>
        Project Zen is a prototype developed for the microsoft hackathon 2023
        event and is provided "as is" without any warranties or guarantees.{" "}
        <br />
        The creators and contributors are not liable for any potential damages
        or issues that may arise from the use of this prototype. Any data or
        information provided by users of Project Zen is solely for demonstration
        purposes and is not stored or used for any commercial or malicious
        intent. By using Project Zen, you agree to these terms and conditions.
        This prototype may contain third-party libraries or components, each
        subject to their respective licenses. <br />
        The creators and contributors of Project Zen do not claim ownership or
        responsibility for these third-party elements. <br />
        <br />
        For any questions or concerns, <b>
          please do not contact us at all.
        </b>{" "}
        Just kidding!
        <br />
        Here you go:{" "}
        <a href="mailto:projectzenhackathon@gmail.com">Leave Feedback.</a>
      </p>
      <Footer />
    </div>
  );
};

export default Legal;
