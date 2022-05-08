import { Link } from "react-router-dom";

interface NavBarProps {}

const NavBar: React.FunctionComponent<NavBarProps> = (props) => {
  return (
    <footer>
      <nav>
        <Link to="/">Home Page</Link>
        <Link to="/films">Films</Link>
      </nav>
    </footer>
  );
};

export default NavBar;
