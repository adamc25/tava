import { useEffect, useRef, useState } from 'react';

const UseFetchData = (dataSource, compareFn, dependencies, initialData) => {
  const [data, setData] = useState(initialData);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const prevDepsRef = useRef();

  useEffect(() => {
  	if (compareFn(prevDepsRef.current, dependencies)) {
  		return
  	}
  	const loadData = async function() {
  		// We need to reset error / loading state to it's initial values on future loads
  		if (error) {
  			setError(null);
  		}
  		if (!loading) {
  			setLoading(true);
  		}
	  	try {
	  		const data = await dataSource(...dependencies);
	  		setData(data)
	  	} catch (error) {
	  		setError(error)
	  	} finally {
	  		prevDepsRef.current = dependencies;
	  		setLoading(false);
	  	}
  	}
  	loadData()
  }, dependencies);

  return { data, loading, error, setData };
};

export default UseFetchData
