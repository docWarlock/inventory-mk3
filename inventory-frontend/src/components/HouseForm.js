import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import apiClient from '../utils/apiClient';

const HouseForm = () => {
    const [house, setHouse] = useState({
        name: '',
        total_area: null,
        unit: ''
    });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const navigate = useNavigate();
    const { id } = useParams();

    useEffect(() => {
        if (id) {
            fetchHouse(id);
        }
    }, [id]);

    const fetchHouse = async (houseId) => {
        try {
            const response = await apiClient.get(`/houses/${houseId}`);
            // Convert backend field names to frontend format
            setHouse({
                name: response.data.name,
                total_area: response.data.total_area || null,
                unit: response.data.unit || ''
            });
        } catch (err) {
            setError('Failed to fetch house');
        }
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        // Convert total_area to number if it's not empty
        let newValue = value;
        if (name === 'total_area' && value !== '') {
            newValue = parseFloat(value);
        }
        setHouse(prev => ({
            ...prev,
            [name]: newValue
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError(null);

        try {
            // Log the data being sent for debugging
            console.log('Sending house data:', house);

            // Prepare data to send - remove total_area if it's null or empty
            let dataToSend = { ...house };
            if (dataToSend.total_area === null || dataToSend.total_area === '') {
                delete dataToSend.total_area;
            }

            if (id) {
                // Update existing house
                await apiClient.put(`/houses/${id}`, dataToSend);
            } else {
                // Create new house
                await apiClient.post('/houses', dataToSend);
            }
            navigate('/houses');
        } catch (err) {
            console.error('Error saving house:', err.response?.data || err.message);
            setError('Failed to save house: ' + (err.response?.data?.message || err.message));
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>{id ? 'Edit House' : 'Add New House'}</h2>
            {error && <div className="error">{error}</div>}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Name:</label>
                    <input
                        type="text"
                        name="name"
                        value={house.name}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Total Area:</label>
                    <input
                        type="number"
                        name="total_area"
                        value={house.total_area}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>Unit:</label>
                    <select
                        name="unit"
                        value={house.unit}
                        onChange={handleChange}
                    >
                        <option value="">Select unit</option>
                        <option value="sqft">Square Feet</option>
                        <option value="sqm">Square Meters</option>
                        <option value="acres">Acres</option>
                        <option value="hectares">Hectares</option>
                    </select>
                </div>
                <button type="submit" disabled={loading}>
                    {loading ? 'Saving...' : 'Save House'}
                </button>
            </form>
        </div>
    );
};

export default HouseForm;
