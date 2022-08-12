import React, {useState, useEffect} from 'react';
import axios from "axios";
import {Button, Form, Container, Modal } from 'react-bootstrap'
import Data from '../api';

const [userData, setUserData] = useState([])
const [refreshData, setRefreshData] = useState(false)

const [changeUserData, setChangeUserData] = useState({"change": false, "id": 0})
const [newUserData, setNewUserData] = useState("")

//gets run at initial loadup
useEffect(() => {
    getAllData();
}, [])

//refreshes the page
if(refreshData){
    setRefreshData(false);
    getAllOrders();
}

const Preview = () => {
  return(<div></div>)
}

function changeUserDataForData() {
  changeUserData.change = false
  var url = "http://localhost:8080/data/update" + changeUserData.id
  
  axios.put(url, newUserData).then(response => {
      console.log(response.status)
      
      if (response.status == 200) {
        setRefreshData(true)
      }
    })
}

function addUserData() {
  setNewUserData(false)
  var url = "http://localhost:8080/data/create"

  axios.post(url, {
    "First Name": newUserData.FirstName,
    "Last Name": newUserData.LastName,
    "Email": newUserData.Email,
    "Phone": newUserData.Phone
  }).then(respons => {
      if (respons.status == 200) {
        setRefreshData(true)
      }
    })
}

function getAllData() {
  var url = "http://localhost:8080/data"
  
  axios.get(url, {
    responsType: 'json'
  }).then(response => {
      if (response.status == 200) {
        setOrders(response.data)
      }
    })
}

function deleteUserData(id) {
  var url = "http://localhost:8080/data/delete/" + id

  axios.delete(url, {
  }).then(response => {
      if (response.status == 200) {
        setRefreshData(true)
      }
    })
}

export default Preview
