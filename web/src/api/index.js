import React, {useState, useEffect} from 'react';
import { Axios } from 'axios';
import {Button, Card, Row, Col} from 'react-bootstrap'

const Data = ({userData, setChangeUserData, deleteUserData}) => {
  return (
    <div>
      <Card>
        <Row>
          <Col>First Name: {userData !== undefined && userData.FirstName}</Col>
          <Col>Last Name: {userData !== undefined && userData.LastName}</Col>
          <Col>Email: {userData !== undefined && userData.Email}</Col>
          <Col>Phone: {userData !== undefined && userData.Phone}</Col>
          <Col><Button onClick={() => changeUserData()}>Edit</Button></Col>
          <Col><Button onClick={() => deleteUserData(userData._id)}>Reset</Button></Col>
        </Row>
      </Card>
    </div>
  )
}

// {firstName, lastName, email, phone, secPhone, address, education, skills, jobs, projects, lists}

function changeUserData() {
  setChangeUserData({
    "change": true,
    "id": userData._id 
  })
}

export default Data
