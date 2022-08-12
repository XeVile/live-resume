import React, { useContext } from "react";
import FormContext from '../../api/ItemContext';

export type Props = {
  field_id: string;
  field_label: string;
  field_place: string;
  field_value: string;
}

const Element: React.FC<Props> = ({
  field_id, field_label, field_place, field_value
}) => {
  const { handleChange } = useContext(FormContext);

  return(
    <div>
      <label htmlFor={field_id}>{field_label ? field_label : "Missing Label"}</label>
      <input
        type="text"
        id={field_id}
        placeholder={field_place ? field_place : ' '}
        onChange={e => handleChange(field_id, e)}/>
    </div>
  );
};

export default Element
