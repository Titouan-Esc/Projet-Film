import { IOneFilm, IFilms } from "../interfaces/App_interface";
import {BsFillHandThumbsUpFill, BsFillHandThumbsDownFill} from "react-icons/bs"

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
              key={index}>
              <div className="content">
                <h2>{value.title}</h2>
                <img src={`https://image.tmdb.org/t/p/w500${value.poster_path}`} alt="poster" />
                <p>{value.overview}</p>
              </div>
              <div className="likes-dislikes">
                <div className="likes">
                  <BsFillHandThumbsUpFill/>
                  <p>{value.likes}</p>
                </div>
                <div className="dislikes">
                  <BsFillHandThumbsDownFill/>
                  <p>{value.dislikes}</p>
                </div>
              </div>
              <div className="comments"></div>
            </section>
          </>
        );
      })}
    </>
  );
};

export default Film;
