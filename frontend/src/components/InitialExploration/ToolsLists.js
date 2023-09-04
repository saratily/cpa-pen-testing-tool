import Tool from "./Tool";

const ToolsList = (props) => {
  return (
    <ul>
      {props.tools.map((tool) => (
        <Tool
          onEditTool={props.onEditTool}
          onDeleteTool={props.onDeleteTool}
          key={tool.ID}
          type={props.type}
          uuid={props.uuid}
          tool={tool} />
      ))}
    </ul>
  );
};

export default ToolsList;
