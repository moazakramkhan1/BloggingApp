import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { HomeComp } from './components/HomeComponent';
import { AddNewBlog } from './components/Add';
import { Details } from './components/blogDetails';
import { UpdateBlog } from './components/Update';
import { DeleteBlog } from './components/Delete';
import { Login } from './components/login';
import { Register } from './components/registerForm';


function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/home" exact element={<HomeComp/>} />
          <Route path="/createblog" element={<AddNewBlog/>} />
          <Route path="/blogdetails/:id" element={<Details/>} />
          <Route path="/updateDetails/:id" element={<UpdateBlog/>} />
          <Route path="/deleteBlog/:id" element={<DeleteBlog/>} />
          <Route path='/login' element={<Login/>} ></Route>
          <Route path='/' element={<Register/>}></Route>
        </Routes>
      </div>
    </Router>
  );
}

export default App;
