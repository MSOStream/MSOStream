import React, { useEffect, useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import CommentView from './components/player/CommentView';
import Player from './components/player/Player';
import client from './global/client';
import userContext from './global/userContext';
import { Layout } from './Layout';
import EditTag from './pages/EditTag';
import Id from './pages/Id';
import Login from './pages/Login';
import PlayerPage from './pages/PlayerPage';
import Register from './pages/Register';
import Upload from './pages/Upload';
import VideoList from './pages/VideoList';
import TagPage from './pages/TagPage';
import EditVideo from './pages/EditVideo';
import Playlist from './pages/Playlist';
import PlaylistPage from './pages/PlaylistPage';

export const App = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  const [userId, setUserId] = useState(null);

  useEffect(() => {
    const getId = () => {
      client
        .get('/user/id', {
          headers: {
            Accept: 'text/html',
            'Content-Type': 'text/plain',
          },
        })
        .then((resp) => {
          setLoggedIn(true);
          setUserId(resp.data);
        })
        .catch(() => null);
    };
    getId();
  });

  const user = {
    loggedIn,
    setLoggedIn,
    userId,
    setUserId,
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route
          path='/'
          element={
            <userContext.Provider value={user}>
              <Layout />
            </userContext.Provider>
          }
        >
          <Route path='player' element={<PlayerPage />}>
            <Route
              index
              element={
                <main>
                  <p>Select a video</p>
                </main>
              }
            />
            <Route
              path=':video_id'
              element={
                <>
                  <Player />
                  <CommentView />
                </>
              }
            />
          </Route>
          <Route path='video/edit/:video_id' element={<EditVideo />} />
          <Route path='videos' element={<VideoList />} />
          <Route path='upload' element={<Upload />} />
          <Route path='playlist' element={<Playlist />} />
          <Route path='playlist'>
            <Route index element={<Playlist />} />
            <Route path=':playlist_id' element={<PlaylistPage />} />
          </Route>
          <Route path='tag'>
            <Route path=':tag_id' element={<TagPage />} />
            <Route path='edit' element={<EditTag />}>
              <Route path=':tag_id' element={<EditTag />} />
            </Route>
          </Route>
          <Route path='login' element={<Login />} />
          <Route path='register' element={<Register />} />
          <Route path='id' element={<Id />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};
