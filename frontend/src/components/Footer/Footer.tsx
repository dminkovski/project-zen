import { styled } from "styled-components";
import { Link } from "react-router-dom";

const StyledFooter = styled.div`
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  text-align: center;
`;

const Footer = () => {
  return (
    <StyledFooter>
      <p>
        <Link to="/about" style={{ color: "#666", marginRight: 5 }}>
          | About Us |
        </Link>
        <Link to="/legal" style={{ marginLeft: 5, color: "#666" }}>
          Legal |
        </Link>
      </p>
      <p> &copy; 2023 Project-Zen. All rights reserved.</p>
    </StyledFooter>
  );
};
export default Footer;
