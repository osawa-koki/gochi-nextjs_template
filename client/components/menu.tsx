import Link from "next/link";
import { useEffect, useState } from "react";
import pages from "../pages";
import { Button } from "react-bootstrap";
import { BsGearFill } from "react-icons/bs";

function Menu() {

  const [current_page, setCurrentPage] = useState<string | null>(null);

  useEffect(() => {
    setCurrentPage(window.location.pathname);
  }, []);

  const PageChanged = () => {
    setCurrentPage(window.location.pathname);
  };

  return (
    <>
      <div id="Menu">
      {
        pages.map((page, index: number) => {
          return (
            <Link key={index} href={page.path} className={`btn ${current_page === page.path ? 'btn-primary' : ''}`} onClick={PageChanged}>
              {page.name}
            </Link>
          )
        })
      }
      </div>
      <div id="ToMenu">
        <Button id="Closer" variant="primary"></Button>
        <BsGearFill id="Opener"></BsGearFill>
      </div>
    </>
  );
};

export default Menu;
