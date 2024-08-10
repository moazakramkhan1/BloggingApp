import React, { useEffect, useState } from "react";
import axios from 'axios';
import { Link, useNavigate} from 'react-router-dom';
import { Navbar } from "./Navbar";
import { FaEdit, FaTrash, FaSpinner } from 'react-icons/fa';
import './styles.css'


export const HomeComp = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const Navigate = useNavigate()
  
  const [name,setName] = useState('')
  const [id,setID] = useState('')
    useEffect(() => {
     (
     async () => {
       const response = await fetch('http://localhost:8000/api/user', {
         headers: { 'Content-Type': 'application/json' },
         credentials: 'include',
       });
       if(response.status === 401){
            setName("")
            setID("")
       }
       else{
       const content = await response.json();
           setName(content.name)
           setID(content.UserID)
       }
     }
     )();
   });

  const fetchData = async () => {
    try {
      setLoading(true);
      const response = await axios.get("http://localhost:8000/bloglist",{withCredentials:true});
      const content = response.data.data;
      setData(content);
      setLoading(false);
    } catch (e) {
      setErrorMessage(e.message);
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);
  if (loading) {
    return (
      <div className="spinner">
        <FaSpinner className="fa-spin" />
      </div>
    );
  }

  if (errorMessage) {
    return (
      <div className="error-message">
        Error Occurred: {errorMessage}
      </div>
    );
  }
  let menu
   if(name === "" && id === ""){
      menu = (
        <>
           <Navbar />
           <div className="not-loggedin-container">
            <div>
              You are not logged in  
            </div>
              </div>
        </>
       
      )
   }
   else{
    menu = (
      <div className="container">
        <div className="header">
          <h1>Blog Application</h1>
          {name && <h2>Welcome {name}</h2>}
        </div>
        <Navbar />
        <button className="newblog-btn" onClick={()=>Navigate("/createblog")}>Create New Blog</button>
        <div className="blog-container">
          {data && data.length > 0 ? data.map((item) => (
            <div className="blog" key={item.id}>
              <Link to={`/blogdetails/${item.id}`}>{item.title}</Link>
              <div className="icons">
                <FaEdit className="icon" onClick={()=>{Navigate(`/updateDetails/${item.id}`)}}/>
                <FaTrash className="icon" onClick={()=>{Navigate(`/deleteBlog/${item.id}`)}}/>
              </div>
              <p>{item.post}</p>
              {item.image && <img src ={`http://localhost:8000${item.image}`} alt={item.title} />}
            </div>
          )) : <p>No blogs available.</p>}
        </div>
      </div>
    );
   }
   return (
        <div>
        {menu}</div>
   )
}
