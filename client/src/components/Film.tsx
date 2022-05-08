import { OneFilm } from "../pages/Films";

interface FilmProps {
  card: OneFilm | undefined;
  setCard: React.Dispatch<React.SetStateAction<OneFilm | undefined>>;
}

const Film: React.FunctionComponent<FilmProps> = () => {
  return <></>;
};

export default Film;
