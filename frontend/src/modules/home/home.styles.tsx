import styled from "styled-components";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
`;
export const Row = styled.div`
  display: flex;
  margin: auto;
  text-align: center;
  flex-direction: row;
  align-items: stretch;
  justify-content: space-evenly;
  gap: 15px;
`;
export const LeftColumn = styled.div`
  flex: 0.5;
`;

export const RightColumn = styled.div`
  flex: 0.5;
`;
