import * as React from "react";
import {
  Table,
  TableVariant,
  TableHeader,
  TableBody,
  IRowData,
  sortable
} from "@patternfly/react-table";

interface ILinkListProps {
  rows: ILink[];
}

export interface ILink {
  role: string;
  name: string;
  address: string;
  deliveries: number;
  rejected: number;
  released: number;
  modified: number;
  presettled: number;
  undelivered: number;
  status?: "creating" | "deleting" | "running";
}

export const LinkList: React.FunctionComponent<ILinkListProps> = ({ rows }) => {
  const toTableCells = (row: ILink) => {
    const tableRow: IRowData = {
      cells: [
        row.role,
        row.name,
        row.address,
        row.deliveries,
        row.rejected,
        row.released,
        row.modified,
        row.presettled,
        row.undelivered
      ],
      originalData: row
    };
    return tableRow;
  };
  const tableRows = rows.map(toTableCells);
  const tableColumns = [
    "Role",
    { title: "Name", transforms: [sortable] },
    "Address",
    "Deliveries",
    { title: "Rejected", transforms: [sortable] },
    { title: "Released", transforms: [sortable] },
    { title: "Modified", transforms: [sortable] },
    { title: "Presettled", transforms: [sortable] },
    { title: "Undelivered", transforms: [sortable] }
  ];

  return (
    <Table
      variant={TableVariant.compact}
      cells={tableColumns}
      rows={tableRows}
      aria-label="links list"
    >
      <TableHeader />
      <TableBody />
    </Table>
  );
};