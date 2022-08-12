import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import {Button, Card, Row, Col} from 'react-bootstrap';
import Element from '../../Element';
import Section from '../../sections'
import FormContext from '../../../api/ItemContext';

const listItems = Section.list.map((item: any) =>
  <Row>
    <Col>
      <Element
        field_id={item.field_id}
        field_label={item.field_label}
        field_place={item.field_place}
        field_value={item.field_value}/>
    </Col>
  </Row>
)

const ListForm: React.FC<{}> = (
  /*{data, setChangeData, deleteData}*/
) => {
  const [allElements, setAllElements] = useState(null);

  const handleChange = (id: string, e: React.FormEvent<HTMLInputElement>) => {
    const newItems = {...Section}

    newItems.list.forEach(field => {
      const { field_id, field_label, field_place, field_value } = field;
      
      if (id === field_id) {
          field.field_value = e.currentTarget.value;
          console.log(field_id, "=", field_value);
        }

      setAllElements(newItems);
      }
    );
  }

  return (
    <Card>
      <FormContext.Provider value={{handleChange}}>
        {listItems}
      </FormContext.Provider>
    </Card>
  )
  
  //function changeInfo() {
  //  setChangeData({
  //    "change": true,
  //    "id"    : data._id
  //  })
  //}
}

export default ListForm
