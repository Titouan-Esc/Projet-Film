import { IOneFilm, IFilms } from "../interfaces/App_interface";

export interface FilmProps {
  film: IOneFilm | undefined;
}

const Film = ({ film }: FilmProps) => {
  return (
    <>
      {film?.map((value: IFilms, index: number, array: IOneFilm) => {
        return (
          <>
            <section
              className="film"
              key={index}
              style={{
                backgroundImage: `url(https://image.tmdb.org/t/p/w500${value.poster_path})`,
              }}
            >
              <div className="content">
                <h2>{value.title}</h2>
                <p>{value.overview}</p>
              </div>
            </section>
            <div className="likes-dislikes"></div>
            <div className="comments"></div>
          </>
        );
      })}
    </>
  );
};

export default Film;
