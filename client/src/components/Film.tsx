import { IOneFilm, IFilms } from "../interfaces/App_interface";

export interface FilmProps {
  film: IOneFilm | undefined;
}

const Film = ({ film }: FilmProps) => {
  return (
    <>
      {film?.map((value: IFilms, index: number, array: IOneFilm) => {
        return (
          <section key={index}>
            <h2>{value.title}</h2>
          </section>
        );
      })}
    </>
  );
};

export default Film;
