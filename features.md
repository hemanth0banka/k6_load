# Features – k6-Inspired Load Tester

## Overview

These are the features which I implemented in my k6 load tester. The tool is inspired by the real k6 load tester and focuses on core load testing capabilities.

---

## 1. Script Generation

- Provides an option to generate load test scripts based on user input
- Scripts follow a JavaScript-based structure similar to k6. A URL is required to create a script.
- Allows configuration of additional optional HTTP request details such as:
  - Request method 
  - Target URL (main)
  - Headers
  - Request body

---

## 2. Local Script Storage

- Generated load test scripts are stored locally on the server
- Maintain scripts with unique identifiers
- Retrieve saved scripts through APIs
- Reuse existing scripts for multiple test runs

---

## 3. Script Download

- Provides a button to download generated load test scripts as JavaScript files

---

## 4. Test Configuration

- Create test configurations using saved scripts
- Configure number of virtual users (VUs)
- Configure test duration
- Link scripts with test execution settings

---

## 5. Load Test Execution

- Provides a run option to execute load tests using selected scripts
- Run requests concurrently
- Simulate multiple virtual users
- Trigger test execution via REST APIs

---

## 6. Metrics Collection

- Collects response time for each request
- Tracks the total number of requests
- Tracks successful and failed requests
- Calculates basic error rate

---

## 7. Test Result Storage

- Stores the results of each test execution
- Maintains execution history
- Retrieves historical test results through APIs

---

## 8. Visualization and Monitoring

- Displays test execution results in the frontend
- Shows response time and throughput charts
- Visualizes test progress during execution
- Displays results of previous test runs

---

## Summary

This k6 Load Tester provides a simple and structured workflow to generate load test scripts, execute them, and visualize results. It focuses on core load testing concepts and serves as a foundation for future improvements.

___________________________________________________________________________


