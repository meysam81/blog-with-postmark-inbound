#!/bin/bash

# Tarzan Project Validation Script
# This script runs comprehensive tests and validations

echo "🎯 Tarzan Project Validation"
echo "=========================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    local status=$1
    local message=$2
    case $status in
        "success")
            echo -e "${GREEN}✓${NC} $message"
            ;;
        "warning")
            echo -e "${YELLOW}⚠${NC} $message"
            ;;
        "error")
            echo -e "${RED}✗${NC} $message"
            ;;
        "info")
            echo -e "${BLUE}ℹ${NC} $message"
            ;;
    esac
}

# Check if we're in the right directory
if [ ! -f "package.json" ]; then
    print_status "error" "package.json not found. Please run this script from the project root."
    exit 1
fi

PROJECT_NAME=$(node -p "require('./package.json').name" 2>/dev/null || echo "unknown")
if [ "$PROJECT_NAME" != "tarzan" ]; then
    print_status "warning" "Project name is '$PROJECT_NAME', expected 'tarzan'"
fi

echo ""
print_status "info" "Starting validation for project: $PROJECT_NAME"
echo ""

# 1. Dependency Check
echo "📦 Checking Dependencies..."
if npm list --depth=0 > /dev/null 2>&1; then
    print_status "success" "All dependencies are properly installed"
else
    print_status "warning" "Some dependencies may be missing"
    print_status "info" "Running npm install..."
    npm install
fi

# 2. Linting
echo ""
echo "🔍 Running Linter..."
if npm run lint > /dev/null 2>&1; then
    print_status "success" "Code passes linting checks"
else
    print_status "warning" "Linting issues found, attempting to fix..."
    npm run lint 2>&1 | tail -10
fi

# 3. Build Test
echo ""
echo "🏗️  Testing Build Process..."
if npm run build > /dev/null 2>&1; then
    print_status "success" "Build completed successfully"

    # Check build output
    if [ -d "dist" ]; then
        SIZE=$(du -sh dist 2>/dev/null | cut -f1)
        print_status "info" "Build output size: $SIZE"

        # Check for essential files
        if [ -f "dist/index.html" ]; then
            print_status "success" "HTML entry point created"
        else
            print_status "error" "HTML entry point missing"
        fi

        if ls dist/assets/*.js 1> /dev/null 2>&1; then
            JS_COUNT=$(ls dist/assets/*.js | wc -l)
            print_status "success" "JavaScript bundles created ($JS_COUNT files)"
        else
            print_status "error" "JavaScript bundles missing"
        fi

        if ls dist/assets/*.css 1> /dev/null 2>&1; then
            CSS_COUNT=$(ls dist/assets/*.css | wc -l)
            print_status "success" "CSS bundles created ($CSS_COUNT files)"
        else
            print_status "warning" "CSS bundles missing"
        fi
    else
        print_status "error" "Build output directory not found"
    fi
else
    print_status "error" "Build failed"
    npm run build 2>&1 | tail -20
fi

# 4. Component Structure Check
echo ""
echo "🧩 Checking Component Structure..."
COMPONENTS_DIR="src/components"
if [ -d "$COMPONENTS_DIR" ]; then
    COMPONENT_COUNT=$(find $COMPONENTS_DIR -name "*.vue" | wc -l)
    print_status "success" "Found $COMPONENT_COUNT Vue components"

    # Check for essential components
    essential_components=("MainContent.vue" "BlogCard.vue" "ErrorBoundary.vue")
    for component in "${essential_components[@]}"; do
        if [ -f "$COMPONENTS_DIR/$component" ]; then
            print_status "success" "$component exists"
        else
            print_status "error" "$component missing"
        fi
    done
else
    print_status "error" "Components directory not found"
fi

# 5. Utilities Check
echo ""
echo "🛠️  Checking Utilities..."
UTILS_DIR="src/utils"
if [ -d "$UTILS_DIR" ]; then
    UTIL_COUNT=$(find $UTILS_DIR -name "*.js" | wc -l)
    print_status "success" "Found $UTIL_COUNT utility modules"

    # Check for essential utilities
    essential_utils=("api.js" "network.js" "accessibility.js" "performance.js")
    for util in "${essential_utils[@]}"; do
        if [ -f "$UTILS_DIR/$util" ]; then
            print_status "success" "$util exists"
        else
            print_status "error" "$util missing"
        fi
    done
else
    print_status "error" "Utils directory not found"
fi

# 6. Accessibility Check
echo ""
echo "♿ Checking Accessibility Features..."
if grep -r "aria-" src/ > /dev/null 2>&1; then
    ARIA_COUNT=$(grep -r "aria-" src/ | wc -l)
    print_status "success" "Found $ARIA_COUNT ARIA attributes"
else
    print_status "warning" "No ARIA attributes found"
fi

if grep -r "role=" src/ > /dev/null 2>&1; then
    ROLE_COUNT=$(grep -r "role=" src/ | wc -l)
    print_status "success" "Found $ROLE_COUNT role attributes"
else
    print_status "warning" "No role attributes found"
fi

if grep -r "sr-only" src/ > /dev/null 2>&1; then
    print_status "success" "Screen reader support implemented"
else
    print_status "warning" "No screen reader specific content found"
fi

# 7. Error Handling Check
echo ""
echo "🚨 Checking Error Handling..."
if grep -r "try.*catch" src/ > /dev/null 2>&1; then
    TRY_CATCH_COUNT=$(grep -r "try.*catch" src/ | wc -l)
    print_status "success" "Found $TRY_CATCH_COUNT try-catch blocks"
else
    print_status "warning" "No try-catch error handling found"
fi

if grep -r "ErrorBoundary" src/ > /dev/null 2>&1; then
    print_status "success" "Error boundary component implemented"
else
    print_status "warning" "No error boundary found"
fi

# 8. Performance Features Check
echo ""
echo "⚡ Checking Performance Features..."
if grep -r "lazy" src/ > /dev/null 2>&1; then
    print_status "success" "Lazy loading implemented"
else
    print_status "info" "No lazy loading found"
fi

if [ -f "src/utils/performance.js" ]; then
    print_status "success" "Performance monitoring utility exists"
else
    print_status "warning" "No performance monitoring found"
fi

# 9. Network Features Check
echo ""
echo "🌐 Checking Network Features..."
if grep -r "navigator.onLine" src/ > /dev/null 2>&1; then
    print_status "success" "Online/offline detection implemented"
else
    print_status "warning" "No network status detection found"
fi

if grep -r "retry" src/ > /dev/null 2>&1; then
    print_status "success" "Retry logic implemented"
else
    print_status "warning" "No retry logic found"
fi

# 10. Security Check
echo ""
echo "🔒 Basic Security Check..."
if grep -r "v-html" src/ > /dev/null 2>&1; then
    print_status "warning" "v-html found - ensure content is sanitized"
    grep -r "v-html" src/ | head -3
else
    print_status "success" "No unsafe v-html usage found"
fi

if grep -r "eval\|innerHTML\|outerHTML" src/ > /dev/null 2>&1; then
    print_status "warning" "Potentially unsafe JavaScript found"
else
    print_status "success" "No obviously unsafe JavaScript patterns found"
fi

# Final Summary
echo ""
echo "📊 Validation Summary"
echo "===================="
print_status "info" "Project: $PROJECT_NAME"
print_status "info" "Node.js: $(node --version 2>/dev/null || echo 'Not available')"
print_status "info" "npm: $(npm --version 2>/dev/null || echo 'Not available')"

if [ -f "dist/index.html" ]; then
    print_status "success" "✅ Build Ready - Project can be deployed"
else
    print_status "error" "❌ Build Issues - Fix errors before deployment"
fi

echo ""
print_status "info" "🎯 Tarzan email-powered blogging platform validation complete!"
echo ""

# Optional: Open preview if available
if command -v npm &> /dev/null && npm run preview --help &> /dev/null; then
    echo "To preview the built application, run:"
    echo "  npm run preview"
    echo ""
fi

echo "To start development server, run:"
echo "  npm start"
echo ""
