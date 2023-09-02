import React, { useState } from 'react';
import styled from 'styled-components';

const Button = styled.button`
  background-color: black;
  color: white;
  font-size: 20px;
  padding: 10px 60px;
  border-radius: 5px;
  margin: 10px 0px;
  cursor: pointer;
  &:disabled {
    color: grey;
    opacity: 0.7;
    cursor: default;
  }
`;

const sayHello = 'Hello';

function ToggleGroup() {
  return (<Button disabled onClick={sayHello}>
    Disabled Button
  </Button>);

}

export default ToggleGroup;


/*
const Button = styled.button`
  /* Same as above * / 
`;
const ButtonToggle = styled(Button)`
  opacity: 0.6;
  ${({ active }) =>
    active &&
    `
    opacity: 1;
  `}
`;

const ButtonGroup = styled.div`
  display: flex;
`;
const types = ['Cash', 'Credit Card', 'Bitcoin'];
function ToggleGroup() {
  const [active, setActive] = useState(types[0]);
  return (
    <ButtonGroup>
      {types.map(type => (
        <ButtonToggle
          key={type}
          active={active === type}
          onClick={() => setActive(type)}
        >
          {type}
        </ButtonToggle>
      ))}
    </ButtonGroup>
  );
}

export default ToggleGroup;

*/

// const FirstTab = () => {
//   return (
    
//     <div className="FirstTab">
//       <p>First Tab!! Hurray!!</p>
//       First tab content will go here
//     </div>
//   );
// };
// export default FirstTab;
