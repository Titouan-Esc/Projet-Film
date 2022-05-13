export type IFilms = {
  id: number;
  poster_path: string;
  overview: string;
  release_date: string;
  title: string;
  popularity: number;
  likes: number;
  dislikes: number;
  comments: string;
};

export type IOneFilm = IFilms[];
