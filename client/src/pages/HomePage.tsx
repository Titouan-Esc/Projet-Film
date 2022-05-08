import NavBar from "../components/NavBar";

interface HomePageProps {}

const HomePage: React.FunctionComponent<HomePageProps> = (props) => {
  return (
    <main>
      <h1>Home Page</h1>
      <NavBar />
    </main>
  );
};

export default HomePage;
