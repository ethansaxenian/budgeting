import React, { useEffect, useState } from "react";
import api from "./apiConfig";
import MonthPage from "./MonthPage";

interface Month {
  id: string;
  startingBalance: number;
  name: number;
  year: number;
}

const App: React.FC = () => {
  const [months, setMonths] = useState<Month[]>([]);
  const [activeMonth, setActiveMonth] = useState<string>();

  useEffect(() => {
    const fetchMonths = async () => {
      try {
        const response = await api.get("/months");
        setMonths(response.data);
        setActiveMonth(response.data[0].id);
      } catch (error) {
        console.error("Error fetching months:", error);
      }
    };

    fetchMonths();
  }, []);


  return (
    <div>
      <nav>
        <ul>
          {months.map((month) => (
            <li
              className={activeMonth === month.id ? "active" : ""}
              onClick={() => setActiveMonth(month.id)}
            >
              {month.id}
            </li>
          ))}
        </ul>
      </nav>
      <MonthPage monthId={activeMonth as string} />
    </div>
  );
};

export default App;
