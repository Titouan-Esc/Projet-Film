import NavBar from "../components/NavBar";

interface OneFilmProps {}

const OneFilm: React.FunctionComponent<OneFilmProps> = () => {
  return (
    <main className="one-film">
      <h1>One Film</h1>
      <NavBar />
    </main>
  );
};

export default OneFilm;
