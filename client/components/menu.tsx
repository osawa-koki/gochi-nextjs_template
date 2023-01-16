import Link from "next/link";
import pages from "../pages";

function Menu() {
  return (
    <div id="Menu">
    {
      pages.map((page, index: number) => {
        return (
          <Link key={index} href={page.path}>
            {page.name}
          </Link>
        )
      })
    }
    </div>
  );
};

export default Menu;
