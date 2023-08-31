import Penetration from "./Penetration";

const PenetrationsList = (props) => {
  return (
    <ul>
      {props.penetrations.map((penetration) => (
        <Penetration
          onEditPenetration={props.onEditPenetration}
          onDeletePenetration={props.onDeletePenetration}
          key={penetration.ID}
          penetration={penetration} />
      ))}
    </ul>
  );
};

export default PenetrationsList;
