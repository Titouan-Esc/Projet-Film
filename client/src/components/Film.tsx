import { IOneFilm, IFilms } from "../interfaces/App_interface";
import {BsFillHandThumbsUpFill, BsFillHandThumbsDownFill} from "react-icons/bs"
import axios, { AxiosResponse } from "axios";
import { useState, MouseEvent, MouseEventHandler } from "react";
import { getMouseEventOptions } from "@testing-library/user-event/dist/utils";

export interface FilmProps {
  film: IOneFilm | undefined;
  getMovies(): any; 
}

const Film = ({ film, getMovies }: FilmProps) => {

  const likes = (idFilm: number) => {
    try {
      axios.put(`${process.env.REACT_APP_API_URL}/movie/like`, {id: idFilm})
      .then((res: AxiosResponse<number>) => {
        getMovies();
      })
    } catch (error) {
      console.log(error);
    }
  }

  const dislikes = (idFilm: number) => {
    try {
      axios.put(`${process.env.REACT_APP_API_URL}/movie/dislike`, {id: idFilm})
      .then((res: AxiosResponse<number>) => {
        getMovies();
      })
    } catch (error) {
      console.log(error);
    }
  }

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
                  <BsFillHandThumbsUpFill size={30} onClick={() => likes(value.id)}/>
                  <p>{value.likes}</p>
                </div>
                <div className="dislikes">
                  <BsFillHandThumbsDownFill size={30} onClick={() => dislikes(value.id)}/>
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
