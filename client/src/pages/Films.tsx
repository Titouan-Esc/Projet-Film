import { useState } from "react";
import Film from "../components/Film";
import NavBar from "../components/NavBar";

interface FilmsProps {}

export interface OneFilm {
  Image: string;
  Title: string;
  Description: string;
}

const Films: React.FunctionComponent<FilmsProps> = () => {
  const [card, setCard] = useState<OneFilm>();
  return (
    <main className="films">
      <h1>Tout les films</h1>
      <Film card={card} setCard={setCard} />
      <NavBar />
    </main>
  );
};

export default Films;
